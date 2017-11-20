package bftsmart

import (
    ab "github.com/hyperledger/fabric/protos/orderer"
    "github.com/hyperledger/fabric/orderer/consensus/bftsmart/persistence"
    "github.com/hyperledger/fabric/orderer/consensus/bftsmart/persistence/file"
    "github.com/hyperledger/fabric/orderer/common/localconfig"
    
    "path/filepath"
    "sync"
)

type MsgDispatch interface {
    
    //This will exec before send message to bftproxy
    BeforeSendMsg(message *ab.BftSmartMessage) *ab.BftSmartMessage
    
    AfterRecvMsg(message *ab.BftSmartMessage)
    
    SetMsgReSender(func (message *ab.BftSmartMessage))
    
    SetConnState(inError bool)
    
}

const (
    MAX_CACHE_MSG = 100
    STORE_FILE_NAME = "message_store"
)


type BFTMsgDisPatch struct {
    //the index of all sended msg
    msgSendIndex    uint64
    //save the sended msg
    sendedCatch     map[uint64]*ab.BftSmartMessage
    //
    msgSender       func (message *ab.BftSmartMessage) error
    //This will be used when the message
    storer          persistence.Storer
    
    //the conn state
    connInErr       bool
    connLock        *sync.Mutex
    
    recvChan        chan *ab.BftSmartMessage
}

func NewBFTMsgDisPatch() *BFTMsgDisPatch{
    
    config := config.Load()
    location :=config.FileLedger.Location
    storeFilePath := location + string(filepath.Separator) + STORE_FILE_NAME
    
    storer := file.NewFileStore(storeFilePath)
    return &BFTMsgDisPatch{
        msgSendIndex:   0,
        sendedCatch:    make(map[uint64]*ab.BftSmartMessage),
        storer:         storer,
        connInErr:      false,
        connLock:       &sync.Mutex{},
    }
}

func (patch *BFTMsgDisPatch) SetMsgReSender(sender func (message *ab.BftSmartMessage) error) {
    patch.msgSender = sender
}

func (patch *BFTMsgDisPatch) BeforeSendMsg(message *ab.BftSmartMessage) *ab.BftSmartMessage {
    
    patch.msgSendIndex++
    message.SendIndex = patch.msgSendIndex
    patch.sendedCatch[patch.msgSendIndex] = message
    
    //save the send msg when the conn is in error
    if patch.connInErr == true {
        if len(patch.sendedCatch) > MAX_CACHE_MSG {
            patch.storer.StoreMessages(patch.sendedCatch)
            patch.sendedCatch = make(map[uint64]*ab.BftSmartMessage)
        }
    }
    
    //TODO deal with the situation when the cache len is very bigger?
    //TODO will this happen when the conn is ok? few possibilites
    
    return message
}

func (patch *BFTMsgDisPatch) AfterRecvMsg(message *ab.BftSmartMessage) {

    //TODO to support
    //remove the msg in sendcache after recv the msg
    msgSendIndex := message.SendIndex
    if _, ok := patch.sendedCatch[msgSendIndex]; ok {
        delete(patch.sendedCatch, msgSendIndex)
    }
    
    
}

//set the conn state
//send the store message when the conn is restore
//store the message when the conn is in error
func (patch *BFTMsgDisPatch) SetConnState(inError bool) {
    patch.connLock.Lock()
    patch.connInErr = inError
    patch.connLock.Unlock()
    
    if patch.connInErr == false {
        patch.reSendMsgWhenConnRestore()
    }
}

func (patch *BFTMsgDisPatch) reSendMsgWhenConnRestore() {
    
    patch.recvChan = make(chan *ab.BftSmartMessage)
    err := patch.storer.GetAllMessage(patch.recvChan)
    
    if err != nil {
    
    }
    
    go func() {
        for {
            msg, ok := <- patch.recvChan
            
            if !ok {
                if len(patch.sendedCatch) > 0 {
                    for _, v := range patch.sendedCatch {
                        patch.msgSender(v)
                    }
                }
                
                break
            }
            
            patch.msgSender(msg)
        }
    }()
}

