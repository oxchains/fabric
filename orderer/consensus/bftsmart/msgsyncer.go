package bftsmart

import (
    cb "github.com/hyperledger/fabric/protos/common"
    "github.com/hyperledger/fabric/orderer/consensus"
    ab "github.com/hyperledger/fabric/protos/orderer"
    "github.com/hyperledger/fabric/protos/utils"
    "github.com/hyperledger/fabric/orderer/common/msgprocessor"
    
    "fmt"
)

type pendingBlock struct {
    pendingMsgs []*syncMsg
}

type syncMsg struct {
    message *cb.Envelope
    //the blockindex of block where the msg store in
    blockIndex uint64
    //the msgoffset
    msgOffset uint64
    //is the last msg
    isOver bool
}

type receiver struct {
    //store the msg ordered by msgoffset
    pendBlock *pendingBlock
    //store the msg whose msgoffset is bigger than offset
    receivedMsg map[uint64]*syncMsg
    //the acquring msgoffset of the next msg which is ordered
    offset uint64
    lastBlockNum uint64
    support      consensus.ConsenterSupport
    syncOver bool
}

func newReceiver(startFrom uint64, lastBlockNum uint64, support consensus.ConsenterSupport) *receiver {
    return &receiver{
        pendBlock: &pendingBlock{
            pendingMsgs: make([]*syncMsg, 0),
        },
        offset: startFrom,
        support: support,
        lastBlockNum: lastBlockNum,
        syncOver: false,
        receivedMsg: make(map[uint64]*syncMsg),
    }
}

func (r *receiver) recvMsg(syncOver bool, msgOffset uint64, blockIndex uint64, env *cb.Envelope) error {
    
    logger.Debugf("Receive a message, offset %d, blockindex %d", msgOffset, blockIndex)
    
    //ignore the msg is already getted
    if msgOffset < r.offset {
        logger.Debugf("Already dealed with the msg offset with %d", msgOffset)
        return nil
    }
    
    //store the msg
    if msgOffset == r.offset {
        recvMsg := &syncMsg{
            message: env,
            blockIndex: blockIndex,
            msgOffset: msgOffset,
            isOver: syncOver,
        }
        r.pendBlock.pendingMsgs = append(r.pendBlock.pendingMsgs, recvMsg)
        r.offset++
        
        err := r.processMsg(recvMsg.message)
        if err != nil {
            logger.Debugf(err.Error())
            return err
        }
        
        //when dealed with a message ,try to write block after
        r.tryToWriteBlock()
        //append the msg in pendingbatch if the has recv the msg right after offset.
        for rcvMsg, ok := r.receivedMsg[r.offset]; ok; rcvMsg, ok = r.receivedMsg[r.offset]{
            r.pendBlock.pendingMsgs = append(r.pendBlock.pendingMsgs, rcvMsg)
            r.processMsg(rcvMsg.message)
            r.tryToWriteBlock()
            delete(r.receivedMsg, r.offset)
            r.offset++
        }
        
    }
    
    //store the msg whose msgoffset is bigger than offset
    if msgOffset > r.offset {
        _, ok := r.receivedMsg[msgOffset]
        if ok {
            logger.Debugf("Already get the msg offset with %d", msgOffset)
            return nil
        }
        
        logger.Debugf("Put the msg %d in cache, current offset is %d", msgOffset, r.offset)
        
        if r.receivedMsg == nil {
            r.receivedMsg = make(map[uint64]*syncMsg)
        }
        
        r.receivedMsg[msgOffset] = &syncMsg{
            message: env,
            blockIndex: blockIndex,
            msgOffset: msgOffset,
            isOver: syncOver,
        }
    }
    
    return nil
}

func (r *receiver) processMsg(env *cb.Envelope) error {
    
    chdr, err := utils.ChannelHeader(env)
    if err != nil {
        return fmt.Errorf("discarding bad config message because of channel header unmarshalling error = %s", err)
    }
    
    class := r.support.ClassifyMsg(chdr)
    switch class {
    case msgprocessor.ConfigMsg:
        if _, _, err := r.support.ProcessConfigMsg(env); err != nil {
        	return fmt.Errorf("disacarding bad config message because = %s", err)
        }

    
    case msgprocessor.NormalMsg:
        if _, err := r.support.ProcessNormalMsg(env); err != nil {
        	return fmt.Errorf("discarding bad normal message because = %s", err)
        }
    
    case msgprocessor.ConfigUpdateMsg:
        if _, _, err := r.support.ProcessConfigUpdateMsg(env); err != nil {
            return fmt.Errorf("disacarding bad config message because = %s", err)
        }
    
    default:
        logger.Panicf("[channel: %s] Unsupported message classification: %v", r.support.ChainID(), class)
    }

    return nil
}

func (r *receiver) tryToWriteBlock() {
    
    logger.Debugf("pendingMsg length is %d, recvMsg not dealed with len is %d, sync is over or not:", len(r.pendBlock.pendingMsgs), len(r.receivedMsg), r.pendBlock.pendingMsgs[len(r.pendBlock.pendingMsgs) - 1].isOver)
    
    batchLen := len(r.pendBlock.pendingMsgs)
    if len(r.pendBlock.pendingMsgs) >= 2 {
        lastMsgBlockIndex := r.pendBlock.pendingMsgs[batchLen - 1].blockIndex
        lastSecondMsgBlockIndex := r.pendBlock.pendingMsgs[batchLen - 2].blockIndex
        
        if lastMsgBlockIndex == lastSecondMsgBlockIndex + 1 {
            tmpBatch := make([]*cb.Envelope, 0)
            in := 0
            var tmpMsgOffset uint64
            
            for index, env := range r.pendBlock.pendingMsgs {
                
                if r.pendBlock.pendingMsgs[index].blockIndex > lastSecondMsgBlockIndex {
                    in = index
                    tmpMsgOffset = r.pendBlock.pendingMsgs[index-1].msgOffset
                    break
                }
                
                tmpBatch = append(tmpBatch, env.message)
            }
            
            logger.Noticef("the index is %d", in)
            
            r.pendBlock.pendingMsgs = r.pendBlock.pendingMsgs[in:]
            
            block := r.support.CreateNextBlock(tmpBatch)
            if lastSecondMsgBlockIndex != block.GetHeader().GetNumber() {
                logger.Warningf("This incidate may be a bug")
            }
            
            encodedLastOffsetPersisted := utils.MarshalOrPanic(&ab.BftSmartMessageMetadata{MsgOffset: tmpMsgOffset, BlockIndex: block.GetHeader().GetNumber()})
            
            if len(tmpBatch) == 1 {
                daLen := len(block.Data.Data)
                logger.Debugf("the length of created block is %d", daLen)
                chdr, _ := utils.ChannelHeader(tmpBatch[0])
                class := r.support.ClassifyMsg(chdr)
                logger.Debugf("to write block")
                logger.Noticef("[Channel %s], the chdr.type is %d",r.support.ChainID(), chdr.Type)
                if class == msgprocessor.ConfigMsg || class == msgprocessor.ConfigUpdateMsg {
                    logger.Debugf("config")
                    r.support.WriteConfigBlock(block, encodedLastOffsetPersisted)
                } else {
                    logger.Debugf("not config")
                    r.support.WriteBlock(block, encodedLastOffsetPersisted)
                }
            } else {
                r.support.WriteBlock(block, encodedLastOffsetPersisted)
            }
            
            r.lastBlockNum++
        }
        
    }
    
    //when last message is a sync over message
    lastPendingMsg := r.pendBlock.pendingMsgs[len(r.pendBlock.pendingMsgs) - 1]
    if len(r.receivedMsg) == 0 && lastPendingMsg.isOver {
        tmpBatch := make([]*cb.Envelope, 0)
        
        for index :=0; index <= len(r.pendBlock.pendingMsgs) - 1; index++ {
            tmpBatch = append(tmpBatch, r.pendBlock.pendingMsgs[index].message)
        }
        
        block := r.support.CreateNextBlock(tmpBatch)
        if lastPendingMsg.blockIndex != block.GetHeader().GetNumber() {
            logger.Warningf("This incidate may be a bug")
        }
        
        encodedLastOffsetPersisted := utils.MarshalOrPanic(&ab.BftSmartMessageMetadata{MsgOffset: lastPendingMsg.msgOffset, BlockIndex: block.GetHeader().GetNumber()})
        
        if len(tmpBatch) == 1 {
            chdr, _ := utils.ChannelHeader(tmpBatch[0])
            class := r.support.ClassifyMsg(chdr)
            if class == msgprocessor.ConfigMsg {
                r.support.WriteConfigBlock(block, encodedLastOffsetPersisted)
            } else {
                r.support.WriteBlock(block, encodedLastOffsetPersisted)
            }
        } else {
            r.support.WriteBlock(block, encodedLastOffsetPersisted)
        }
        
        r.lastBlockNum++
        r.syncOver = true
    }
}




