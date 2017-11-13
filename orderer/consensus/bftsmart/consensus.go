/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bftsmart

import (
	"time"
	"os"
	"sync"
	"fmt"
	
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/op/go-logging"
	"github.com/hyperledger/fabric/orderer/consensus"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/hyperledger/fabric/orderer/common/msgprocessor"
	
)

var logger = logging.MustGetLogger("orderer/bftsmart")
var poolsize uint = 0
var poolindex uint = 0
var recvport uint = 0

type consenter struct{
}

type chain struct {
	connection *connection
	support      consensus.ConsenterSupport
	batchTimeout time.Duration
	
	lastMsgOffset uint64
	lastCutBlockNumber  uint64
	sender string
	
	sendChan     chan *ab.BftSmartMessage
	recvChan     chan *ab.BftSmartProxy
	exitChan     chan struct{}
	timer         <- chan time.Time
	
	inSync       bool
	inSyncLock   *sync.Mutex
	syncMsgOffset uint64
	
	msgSyncer *receiver
}

const (
	DEFAULT_OFFSET = 0
)

var conn *connection
var sender string

func New(size uint, send uint, recv uint) consensus.Consenter {
	poolsize = size
	recvport = recv
	sender = getOrdererIdentity()
	
	var err error
	conn, err = newConnection()
	if err != nil {
		logger.Panic(err.Error())
	}
	
	return &consenter{
	}
}

func (solo *consenter) HandleChain(support consensus.ConsenterSupport, metadata *cb.Metadata) (consensus.Chain, error) {
	lastMsgOffset := getLastMsgOffset(metadata.Value, support.ChainID())
	return newChain(support, conn, lastMsgOffset), nil
}

func newChain(support consensus.ConsenterSupport, conn *connection, lastMsgOffset uint64) *chain {
	
	lastCutBlockNumber := getLastCutBlockNumber(support.Height())
	return &chain{
		batchTimeout: support.SharedConfig().BatchTimeout(),
		support:      support,
		recvChan:     make(chan *ab.BftSmartProxy),
		sendChan:     make(chan *ab.BftSmartMessage),
		exitChan:     make(chan struct{}),
		connection:   conn,
		timer:        nil,
		lastMsgOffset: lastMsgOffset,
		lastCutBlockNumber: lastCutBlockNumber,
		sender: sender,
		inSync: false,
		inSyncLock: &sync.Mutex{},
		msgSyncer: &receiver{},
	}
}

func (ch *chain) Start() {
	ch.connection.channels[ch.support.ChainID()] = channel{
		ch.support.ChainID(),
		*ch,
	}
	
	go ch.sendBFTMessage()
	go ch.processMessagesToBlocks()
	
	ch.sendConnectMessage()
}

func (ch *chain) Halt() {

	select {
	case <-ch.exitChan:
		// Allow multiple halts without panic
	default:
		close(ch.exitChan)
	}
}

// Configure accepts configuration update messages for ordering (JCS: for the moment, this orderer doe not support this feature)
func (ch *chain) Configure(configMsg *cb.Envelope, configSeq uint64) error {
	
	config, err := utils.Marshal(configMsg)
	if err != nil {
		logger.Debugf("cannot enqueue, unable to marshal config because = %s", err.Error())
	}
	
    msg := ch.newConfigMessage(config, configSeq)
	
	select {
	case ch.sendChan <- msg:
		return nil
	case <-ch.exitChan:
		return fmt.Errorf("Exiting")
	}
    return nil
}

// Errored only closes on exit
func (ch *chain) Errored() <-chan struct{} {
	return ch.exitChan
}

// Order accepts a message and returns true on acceptance, or false on shutdown
func (ch *chain) Order(env *cb.Envelope, configSeq uint64) error {
	
	payload, err := utils.Marshal(env)
	if err != nil {
		return fmt.Errorf("cannot enqueue, unable to marshal config because = %s", err)
	}
	
	msg := ch.newNormalMessage(payload, configSeq)
	
    select {
    case ch.sendChan <- msg:
    	return nil
    case <-ch.exitChan:
	    return fmt.Errorf("Exiting")
    }
}

func (ch *chain) sendBFTMessage() {
	for {
		select {
		case mes := <- ch.sendChan:
			
			poolindex := poolindex % poolsize
			_, err := ch.connection.sendEnvToBFTProxy(mes, poolindex)
			poolindex++
			
			if err != nil {
				logger.Warningf(err.Error())
				continue
			}
		}
	}
}

func (ch *chain) sendConnectMessage() {
	connectMsg := ch.newConnectMessage(nil, ch.lastMsgOffset)
	ch.sendChan <- connectMsg
}

func (ch *chain) sendTimeToCut(timeToCutBlockNumber uint64, timer *<-chan time.Time) error {
	logger.Debugf("[channel: %s] Time-to-cut block %d timer expired", ch.support.ChainID(), timeToCutBlockNumber)
	
	*timer = nil
	msg := ch.newTimerToCutMessage(timeToCutBlockNumber)
	ch.sendChan <- msg
	return nil
}

func (ch *chain) newConnectMessage(payload []byte, msgOffset uint64) *ab.BftSmartMessage{
	return &ab.BftSmartMessage{
	    Type: &ab.BftSmartMessage_Connect{
	        Connect: &ab.BftSmartMessageConnect{
	           Payload: payload,
	           StartFrom: msgOffset,
	        },
	    },
	    ChannelId: ch.support.ChainID(),
	    Sender: ch.sender,
	}
}

func (ch *chain) newNormalMessage(payload []byte, configSeq uint64) *ab.BftSmartMessage{
    return &ab.BftSmartMessage{
    	Type: &ab.BftSmartMessage_Regular{
    	    Regular: &ab.BftSmartMessageRegular{
    	        Payload: payload,
    	        ConfigSeq: configSeq,
    	        Class: ab.BftSmartMessageRegular_NORMAL,
	        },
	    },
	    ChannelId: ch.support.ChainID(),
	    Sender: ch.sender,
    }
}

func (ch *chain) newConfigMessage(config []byte, configSeq uint64) *ab.BftSmartMessage {
	return &ab.BftSmartMessage{
	    Type: &ab.BftSmartMessage_Regular{
	        Regular: &ab.BftSmartMessageRegular{
		        Payload: config,
		        ConfigSeq: configSeq,
		        Class: ab.BftSmartMessageRegular_CONFIG,
	        },
	    },
	    ChannelId: ch.support.ChainID(),
	    Sender: ch.sender,
	}
}

func (ch *chain) newTimerToCutMessage(blockNumber uint64) *ab.BftSmartMessage {
    return &ab.BftSmartMessage{
    	Type: &ab.BftSmartMessage_TimeToCut{
    	    TimeToCut: &ab.BftSmartMessageTimeToCut{
    	        BlockNumber: blockNumber,
	        },
	    },
	    ChannelId: ch.support.ChainID(),
	    Sender: ch.sender,
    }
}

func (ch *chain) NewSynchronizeResponseMessage(payload []byte, msgOffset uint64, blockIndex uint64, isOver bool) *ab.BftSmartMessage {
    return &ab.BftSmartMessage{
        Type: &ab.BftSmartMessage_Synchronize{
            Synchronize: &ab.BftSmartMessageSynchronize{
                Payload: payload,
                MsgOffset: msgOffset,
                Class: ab.BftSmartMessageSynchronize_RESPONSE,
                BlockIndex: blockIndex,
                IsOver: isOver,
            },
        },
        ChannelId: ch.support.ChainID(),
        Sender: ch.sender,
    }
}

func (ch *chain) newSyncronizeRequestMessage(msgStartFrom uint64, blockStartFrom uint64) *ab.BftSmartMessage {
	return &ab.BftSmartMessage{
		Type: &ab.BftSmartMessage_Synchronize{
			Synchronize: &ab.BftSmartMessageSynchronize{
				Payload: nil,
				SyncStartFrom: msgStartFrom,
				BlockIndex: blockStartFrom,
				Class: ab.BftSmartMessageSynchronize_REQUEST,
			},
		},
		ChannelId: ch.support.ChainID(),
		Sender: ch.sender,
	}
}

func (ch *chain) processMessagesToBlocks() ([]uint64, error) {
	
	for {
		select {
		case in := <- ch.recvChan:
			msg := in.BfgMsg
			msgOffset := in.Offset
			needSync := in.GetNeedSync()
			
			switch msg.Type.(type) {
			case *ab.BftSmartMessage_TimeToCut:
				if ch.inSync {
					logger.Debugf("[Orderer %s,Channel %s]In sync to others, ignore the timeout message", ch.sender, ch.support.ChainID())
					return nil, nil
				}
			    ch.processTimeToCut(msg.GetTimeToCut(), msgOffset)
			
			case *ab.BftSmartMessage_Regular:
				if ch.inSync {
					logger.Debugf("[Orderer %s,Channel %s]In sync to others, ignore the normal message", ch.sender, ch.support.ChainID())
					return nil, nil
				}
				if err := ch.processRegular(msg.GetRegular(), msgOffset); err != nil {
			        logger.Warningf("[channel: %s] Error when processing incoming message of type REGULAR = %s", ch.support.ChainID(), err)
			    }
			case *ab.BftSmartMessage_Connect:
				_ = ch.processConnectMessage(msg, needSync)
			
			case *ab.BftSmartMessage_Synchronize:
				syncSender := msg.GetSender()
				if err := ch.processSynchronize(msg.GetSynchronize(), msgOffset, syncSender); err != nil {
					return nil, err
				}
			 
			}
			
		case <- ch.timer:
		    if err := ch.sendTimeToCut(ch.lastCutBlockNumber+1, &ch.timer); err != nil {
			    logger.Errorf("[channel: %s] cannot post time-to-cut message = %s", ch.support.ChainID(), err)
		    }
		}
	}
}

func (ch *chain) processConnectMessage(message *ab.BftSmartMessage, needSync bool) error {
	
	if ch.sender != message.GetSender() {
		logger.Debugf("[Channel %s]This is a connect message from %s, just ignore it", ch.support.ChainID(), message.GetSender())
		return nil
	} else {
		logger.Debugf("[Channel %s]This ia connect message send by itself", ch.support.ChainID())
	}
	
	if needSync {
		ch.inSyncLock.Lock()
		ch.inSync = true
		ch.inSyncLock.Unlock()
		
		syncRequest := ch.newSyncronizeRequestMessage(ch.lastMsgOffset, ch.lastCutBlockNumber)
		startFrom := ch.lastMsgOffset + 1
		ch.msgSyncer = newReceiver(startFrom, ch.lastCutBlockNumber, ch.support)
		ch.sendChan <- syncRequest
	} else {
		ch.inSyncLock.Lock()
		ch.inSync = false
		ch.inSyncLock.Unlock()
		ch.msgSyncer = nil
	}
	
	return nil
}

func (ch *chain) processRegular(regularMessage *ab.BftSmartMessageRegular, receivedOffset uint64) error {
	
	commitNormalMsg := func(message *cb.Envelope) {
		batches, pending := ch.support.BlockCutter().Ordered(message)
		logger.Debugf("[channel: %s] Ordering results: items in batch = %d, pending = %v", ch.support.ChainID(), len(batches), pending)
		if len(batches) == 0 && ch.timer == nil {
			ch.timer = time.After(ch.support.SharedConfig().BatchTimeout())
			logger.Debugf("[channel: %s] Just began %s batch timer", ch.support.ChainID(), ch.support.SharedConfig().BatchTimeout().String())
			return
		}
		
		//offset if the normal offset
		//pending: indicate if there has any message to deal with
		//the length of batches only have three options 0,1,2;
		//batches:0, pending:true, will not create block
		//batches:1, pending:false,will create a block include a message.
		//batches:1, pending:true,will create a block and current message while in pending(not create block).
		//batches:2, pending:
		
		offset := receivedOffset
		if pending || len(batches) == 2 {
			offset--
		}
		
		for _, batch := range batches {
			block := ch.support.CreateNextBlock(batch)
			encodedLastOffsetPersisted := utils.MarshalOrPanic(&ab.BftSmartMessageMetadata{MsgOffset: offset, BlockIndex: block.GetHeader().GetNumber()})
			ch.support.WriteBlock(block, encodedLastOffsetPersisted)
			ch.lastCutBlockNumber++
			offset++
			logger.Debugf("[channel: %s] Batch filled, just cut block %d - last persisted offset is now %d", ch.support.ChainID(), ch.lastCutBlockNumber, offset)
		}
		
		if len(batches) > 0 {
			ch.timer = nil
		}
	}
	
	commitConfigMsg := func(message *cb.Envelope) {
		logger.Debugf("[channel: %s] Received config message", ch.support.ChainID())
		batch := ch.support.BlockCutter().Cut()
		
		if batch != nil {
			logger.Debugf("[channel: %s] Cut pending messages into block", ch.support.ChainID())
			block := ch.support.CreateNextBlock(batch)
			encodedLastOffsetPersisted := utils.MarshalOrPanic(&ab.BftSmartMessageMetadata{MsgOffset: receivedOffset - 1, BlockIndex: block.GetHeader().GetNumber()})
			ch.support.WriteBlock(block, encodedLastOffsetPersisted)
			ch.lastCutBlockNumber++
		}
		
		logger.Debugf("[channel: %s] Creating isolated block for config message, and the offset is %d", ch.support.ChainID(), receivedOffset)
		block := ch.support.CreateNextBlock([]*cb.Envelope{message})
		encodedLastOffsetPersisted := utils.MarshalOrPanic(&ab.BftSmartMessageMetadata{MsgOffset: receivedOffset, BlockIndex: block.GetHeader().GetNumber()})
		ch.support.WriteConfigBlock(block, encodedLastOffsetPersisted)
		ch.lastCutBlockNumber++
		ch.timer = nil
	}
	
	seq := ch.support.Sequence()
	
	env := &cb.Envelope{}
	if err := proto.Unmarshal(regularMessage.Payload, env); err != nil {
		// This shouldn't happen, it should be filtered at ingress
		return fmt.Errorf("failed to unmarshal payload of regular message because = %s", err)
	}
	
	logger.Debugf("[channel: %s] Processing regular bft message of type %s", ch.support.ChainID(), regularMessage.Class.String())
	
	switch regularMessage.Class {
	case ab.BftSmartMessageRegular_UNKNOWN:
		// Received regular message of type UNKNOWN, indicating it's from v1.0.x orderer
		chdr, err := utils.ChannelHeader(env)
		if err != nil {
			return fmt.Errorf("discarding bad config message because of channel header unmarshalling error = %s", err)
		}
		
		class := ch.support.ClassifyMsg(chdr)
		switch class {
		case msgprocessor.ConfigMsg:
			if _, _, err := ch.support.ProcessConfigMsg(env); err != nil {
				return fmt.Errorf("discarding bad config message because = %s", err)
			}
			
			logger.Noticef("[Channel %s], process the config msg",ch.support.ChainID())
			
			commitConfigMsg(env)
		
		case msgprocessor.NormalMsg:
			logger.Noticef("[Channel %s], process the normal msg",ch.support.ChainID())
			if _, err := ch.support.ProcessNormalMsg(env); err != nil {
				return fmt.Errorf("discarding bad normal message because = %s", err)
			}
			
			commitNormalMsg(env)
		
		case msgprocessor.ConfigUpdateMsg:
			logger.Noticef("[Channel %s], process the configupdatemsg msg",ch.support.ChainID())
			return fmt.Errorf("not expecting message of type ConfigUpdate")
		
		default:
			logger.Panicf("[channel: %s] Unsupported message classification: %v", ch.support.ChainID(), class)
		}
	
	case ab.BftSmartMessageRegular_NORMAL:
		if regularMessage.ConfigSeq < seq {
			logger.Debugf("[channel: %s] Config sequence has advanced since this normal message being validated, re-validating", ch.support.ChainID())
			if _, err := ch.support.ProcessNormalMsg(env); err != nil {
				return fmt.Errorf("discarding bad normal message because = %s", err)
			}
			
			// TODO re-submit stale normal message via `Order`, instead of discarding it immediately. Fix this as part of FAB-5720
			return fmt.Errorf("discarding stale normal message because config seq has advanced")
		}
		
		commitNormalMsg(env)
	
	case ab.BftSmartMessageRegular_CONFIG:
		if regularMessage.ConfigSeq < seq {
			logger.Debugf("[channel: %s] Config sequence has advanced since this config message being validated, re-validating", ch.support.ChainID())
			_, _, err := ch.support.ProcessConfigMsg(env)
			if err != nil {
				return fmt.Errorf("rejecting config message because = %s", err)
			}
			
			return fmt.Errorf("discarding stale config message because config seq has advanced")
		}
		
		commitConfigMsg(env)
	
	default:
		return fmt.Errorf("unsupported regular kafka message type: %v", regularMessage.Class.String())
	}
	
	return nil
}

func (ch *chain) processSynchronize(synchronizeMessage *ab.BftSmartMessageSynchronize, receivedOffset uint64, syncSender string) error {
	
	switch synchronizeMessage.Class {
	case ab.BftSmartMessageSynchronize_UNKNOWN:
	    return fmt.Errorf("[Channel %s]Unknow synchronize message", ch.support.ChainID())
	    
	case ab.BftSmartMessageSynchronize_REQUEST:
		if ch.inSync {
			logger.Debugf("[Orderer %s,Channel %s]In sync to others, ignore the sync request message", ch.sender, ch.support.ChainID())
			return nil
		}
		
		if ch.sender == syncSender {
			logger.Debugf("[Channel %s]This is a sync request send by itself,just ignore it", ch.support.ChainID())
			return nil
		}
		
		msgStart := synchronizeMessage.GetSyncStartFrom()
		blockStartFrom := synchronizeMessage.GetBlockIndex()
		ch.processSynchronizeRequestMessage(msgStart, blockStartFrom)
		return nil
	
	case ab.BftSmartMessageSynchronize_RESPONSE:
		
		if ch.sender == syncSender {
			logger.Debugf("[Channel %s]This is sync response send by itself", ch.support.ChainID())
		    return nil
		}
		
		if !ch.inSync {
			logger.Debugf("[Channel %s]Not in sync state, just ignore it")
			return nil
		}
		
	    ch.processSynchronizeResponseMessage(synchronizeMessage, receivedOffset)
	}

	return nil
}

func (ch *chain) processTimeToCut(tt *ab.BftSmartMessageTimeToCut, receivedOffset uint64) error {
	blockNum := tt.GetBlockNumber()
	logger.Debugf("[Channel %s] The blockNum of TimeToCut is %d, and lastCutBlockNumber is %d",ch.support.ChainID(), blockNum, ch.lastCutBlockNumber)
	
	if blockNum == ch.lastCutBlockNumber + 1 {
		ch.timer = nil
		batch := ch.support.BlockCutter().Cut()
		if len(batch) == 0 {
			return fmt.Errorf("Batch timer expired with no pending requests, this might indicate a bug")
		}
		
		block := ch.support.CreateNextBlock(batch)
		encodedLastOffsetPersisted := utils.MarshalOrPanic(&ab.BftSmartMessageMetadata{MsgOffset: receivedOffset, BlockIndex: block.GetHeader().GetNumber()})
		ch.support.WriteBlock(block, encodedLastOffsetPersisted)
		ch.lastCutBlockNumber++
	}
	
	return nil
}

//process Synchronize Message from proxy
//get the message requested before and send it to proxy
func (ch *chain) processSynchronizeRequestMessage(msgStartFrom uint64, blockStartFrom uint64) {
	
	//get the next block
	msgStartFrom++
	blockStartFrom++
	
	logger.Debugf("[Channel %s]Try to deal with the request to sync to other orderers, block startFrom %d, msgStartFrom", ch.support.ChainID(), blockStartFrom, msgStartFrom)
    ch.syncMsgOffset = msgStartFrom
	
    ch.sendSyncMessage(msgStartFrom, blockStartFrom)
}

func (ch *chain) sendSyncMessage(msgStartFrom uint64, blockStartFrom uint64) {

    endBlockIndex := ch.support.Height() - 1
	lastBlockMsgStartIndex := msgStartFrom
	
    for blockIndex := blockStartFrom; blockIndex <= endBlockIndex; blockIndex++ {
    	
    	block := ch.support.GetBlock(blockIndex)
	    
    	bftSmartMetaData := &ab.BftSmartMessageMetadata{}
	    metadata, err := utils.GetMetadataFromBlock(block, cb.BlockMetadataIndex_ORDERER)
	    err = proto.Unmarshal(metadata.Value, bftSmartMetaData)
	    if err != nil {
		    logger.Error(err.Error())
	    }
	
	    logger.Debugf("message start Index is %d, block message index is %d", msgStartFrom, bftSmartMetaData.BlockIndex)
	
	    if bftSmartMetaData.MsgOffset < msgStartFrom {
		    continue
	    }
	
	    messages, err := utils.ExtractAllEnvelopes(block)
	    if err != nil {
		    logger.Error(err.Error())
	    }
	
	    for msgIndex := uint64(0); msgIndex < uint64(len(messages)); msgIndex++ {
		    tmpMetaData := bftSmartMetaData
		    tmpMetaData.MsgOffset = lastBlockMsgStartIndex + msgIndex
		
		    //Err never occur
		    payload, _ := proto.Marshal(messages[msgIndex])
		    
		    over := false
		    if blockIndex == endBlockIndex && msgIndex == uint64(len(messages)) - 1 {
			    over = true
		    }
		
		    logger.Debugf("[Channel %s]Send sync response to other orderers, msgoffset %d, blockindex %s", ch.support.ChainID(), tmpMetaData.MsgOffset, tmpMetaData.BlockIndex)
		    
		    ch.sendChan <- ch.NewSynchronizeResponseMessage(payload, tmpMetaData.MsgOffset, tmpMetaData.BlockIndex, over)
		
	    }
	    lastBlockMsgStartIndex++
    }
}



func (ch *chain) processSynchronizeResponseMessage(synchronizeMessage *ab.BftSmartMessageSynchronize, receivedOffset uint64) error {
	
	logger.Debugf("[Channel %s]process Syncronize Response message", ch.support.ChainID())
	receivedOffset = synchronizeMessage.GetMsgOffset()
	blockIndex := synchronizeMessage.GetBlockIndex()
	syncOver := synchronizeMessage.GetIsOver()
	
	env := &cb.Envelope{}
	if err := proto.Unmarshal(synchronizeMessage.Payload, env); err != nil {
		// This shouldn't happen, it should be filtered at ingress
		return fmt.Errorf("failed to unmarshal payload of regular message because = %s", err)
	}
	
	err := ch.msgSyncer.recvMsg(syncOver, receivedOffset, blockIndex, env)
	
	if err != nil {
		return err
	}
	
	if ch.msgSyncer.syncOver {
		ch.lastMsgOffset = ch.msgSyncer.offset - 1
		ch.lastCutBlockNumber = ch.msgSyncer.lastBlockNum
		//TODO to promise this will after
		defer ch.sendConnectMessage()
	}
	
	return nil
}

func getLastMsgOffset(metadataValue []byte, chainId string) uint64 {
	
    if metadataValue != nil {
        bftSmartMetaData := &ab.BftSmartMessageMetadata{}
        if err := proto.Unmarshal(metadataValue, bftSmartMetaData); err != nil {
	        logger.Panicf("[channel: %s] Ledger may be corrupted:"+
	        "cannot unmarshal orderer metadata in most recent block", chainId)
        }
        return bftSmartMetaData.MsgOffset
    }
    
    return DEFAULT_OFFSET
}

func getLastCutBlockNumber(blockchainHeight uint64) uint64 {
	return blockchainHeight - 1
}

//used in bftmessage to recognize orderer
func getOrdererIdentity() string {
	timeUTC := time.Now().UTC().String()
	host, err := os.Hostname()
	
	if err != nil {
		return timeUTC
	} else {
		return host + timeUTC
	}
}



