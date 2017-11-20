package persistence

import (
    ab "github.com/hyperledger/fabric/protos/orderer"
)

type Storer interface {
    
    StoreMessages(msgs map[uint64]*ab.BftSmartMessage) error
    
    GetAllMessage(recv chan *ab.BftSmartMessage) error
    
}
