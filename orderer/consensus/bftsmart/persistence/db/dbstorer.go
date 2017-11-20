package db

import (
    ab "github.com/hyperledger/fabric/protos/orderer"
    "github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
    "github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
    
    "github.com/golang/protobuf/proto"
    "github.com/op/go-logging"
    "strings"
)

var logger = logging.MustGetLogger("bftsmart/persistence/db")

const (
    DB_NAME = "dbstore"
    NS = "orderer_cache_messages"
    KEYS = "keys"
    
    SEPARATOR = ","
)

type DBStore struct {
    db privacyenabledstate.DB
    
}

func NewDBStore() (*DBStore, error) {
    dbProvider, err := privacyenabledstate.NewCommonStorageDBProvider()
    if err != nil {
        return nil, err
    }
    
    db, err := dbProvider.GetDBHandle(DB_NAME)
    if err != nil {
        return nil, err
    }
    
    dbstore := &DBStore{}
    dbstore.db = db
    
    return dbstore, nil
}

func (dbstore *DBStore) StoreMessages(msgs map[uint64]*ab.BftSmartMessage) error {
    
    keysV, err := dbstore.db.GetState(NS, KEYS)
    if err != nil {
        return err
    }
    
    var keysString string
    if keysV == nil || len(keysV.Value) == 0 {
        keysString = ""
    } else {
        keysString = string(keysV.Value)
    }
    
    batch := statedb.NewUpdateBatch()
    
    for key, msg := range msgs {
        msgBytes, err := proto.Marshal(msg)
        if err != nil {
            continue
        }
        
        batch.Put(NS, string(key), msgBytes, nil)
        keysString = keysString + SEPARATOR + string(key)
    }
    dbstore.db.ApplyUpdates(batch, nil)
    
    return nil
}

func (dbstore *DBStore) GetAllMessage(recv chan *ab.BftSmartMessage) error {
    
    keysV, err := dbstore.db.GetState(NS, KEYS)
    if err != nil {
        close(recv)
        return err
    }
    
    if keysV == nil || len(keysV.Value) == 0 {
        close(recv)
        return nil
    }
    
    keysS := string(keysV.Value)
    keys := strings.Split(keysS, SEPARATOR)
    
    for _, key := range keys {
        value, err := dbstore.db.GetState(NS, key)
        if err != nil {
            continue
        }
        
        var msg ab.BftSmartMessage
        err = proto.Unmarshal(value.Value, &msg)
        if err != nil {
            continue
        }
    
        recv <- &msg
    }
    
    return nil
}
