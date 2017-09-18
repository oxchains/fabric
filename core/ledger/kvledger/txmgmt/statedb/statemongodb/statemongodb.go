package statemongodb

import (
	"mgo-2"
	"fmt"
	"unicode/utf8"
	"encoding/json"
	
	"github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/statedb/mongodbhelper"
	"github.com/oxchains/fabric/common/flogging"
	"github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/version"
	"sync"
)


var logger = flogging.MustGetLogger("statemongodb")

var savePointKey = "statedb_savepoint"
var savePointNs = "savepoint"
//分页使用
var queryskip = 0

type VersionedDBProvider struct {
	session *mgo.Session
	databases map[string]*VersionedDB
	mux sync.Mutex
}

type VersionedDB struct {
	mongoDB *mongodbhelper.MongoDB
	dbName string
}

func NewVersionedDBProvider() (*VersionedDBProvider, error) {
	mongodbConf := mongodbhelper.GetMongoDBConf()
	dialInfo, err := mgo.ParseURL(mongodbConf.Url)
	if err != nil {
	    return nil, err
	}
	dialInfo.Timeout = mongodbConf.RequestTimeout
	mgoSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}
	
	return &VersionedDBProvider{
		session: mgoSession,
		databases: make(map[string]*VersionedDB),
		mux: sync.Mutex{},
	}, nil
}

func (provider *VersionedDBProvider) GetDBHandle(dbName string) (statedb.VersionedDB, error){
	provider.mux.Lock()
	defer provider.mux.Unlock()
	
	db, ok := provider.databases[dbName]
	if ok {
		return db, nil
	}
	
	vdr, err := newVersionDB(provider.session, dbName)
	if err != nil {
		return nil, err
	}
	
	return vdr, nil
}

func (provider *VersionedDBProvider) Close() {
	provider.session.Close()
}

func (vdb *VersionedDB) Close() {
	vdb.mongoDB.Close()
}

func (vdb *VersionedDB) Open() error {
    return nil
}

func (vdb *VersionedDB) GetState(namespace string, key string)(*statedb.VersionedValue, error){
	doc , err := vdb.mongoDB.GetDoc(namespace, key)
	if err != nil{
		return nil, err
	}
	if doc == nil{
		return nil, nil
	}
	
	versionedV := statedb.VersionedValue{Value: nil, Version:& doc.Version,}
	if doc.Value == nil && doc.Attachments != nil {
		versionedV.Value = doc.Attachments.AttachmentBytes
	}else{
		tempbyte , _ := json.Marshal(doc.Value)
		versionedV.Value = tempbyte
	}
	
	return &versionedV, nil
}

func (vdb *VersionedDB) GetStateMultipleKeys(namespace string, keys []string) ([]*statedb.VersionedValue, error) {
	vals := make([]*statedb.VersionedValue, len(keys))
	for i, key := range keys {
		val, err := vdb.GetState(namespace, key)
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}
	return vals, nil
}

func (vdb *VersionedDB) ValidateKey(key string) error {
	if !utf8.ValidString(key) {
		return fmt.Errorf("Key should be a valid utf8 string: [%x]", key)
	}
	return nil
}

func (vdb *VersionedDB) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (statedb.ResultsIterator, error) {
	querylimit := vdb.mongoDB.Conf.QueryLimit
	dbItr := vdb.mongoDB.GetIterator(namespace, startKey, endKey, querylimit, queryskip)
	return newKVScanner(dbItr, namespace), nil
}

func (vdb *VersionedDB) ExecuteQuery(namespace ,query string)(statedb.ResultsIterator, error) {
	
	queryByte := []byte(query)
	if !isJson(queryByte) {
	    return nil, fmt.Errorf("the query is not a json : %s", query)
	}

	queryBson, err := mongodbhelper.GetQueryBson(namespace, query)
	if err != nil {
	    return nil, err
	}
	
	result, err := vdb.mongoDB.QueryDocuments(queryBson)
	if err != nil {
	    return nil, err
	}
	
	return newKVScanner(result, namespace), nil
}

func(vdb VersionedDB) ApplyUpdates(batch *statedb.UpdateBatch, height *version.Height) error {
    namespaces := batch.GetUpdatedNamespaces()
	var out interface{}
	var err error
	for _, ns := range namespaces {
	    updates := batch.GetUpdates(ns)
	    for k, vv := range updates {
		    logger.Infof("Channel [%s]: Applying key=[%#v]", vdb.dbName, k)
		    if vv.Value == nil {
			    vdb.mongoDB.Delete(ns, k)
		    } else {
			    doc := mongodbhelper.MongodbDoc{
				    Key:         k,
				    Version:     *vv.Version,
				    ChaincodeId: ns,
			    }
			
			    if !isJson(vv.Value) {
				    logger.Infof("Not a json, write it to the attachment ")
				    attachmentByte := mongodbhelper.Attachment{
					    AttachmentBytes: vv.Value,
				    }
				    doc.Attachments = &attachmentByte
				    doc.Value = nil
			    } else {
				    err = json.Unmarshal(vv.Value, &out)
				    if err != nil {
					    logger.Errorf("Error rises while unmarshal the vv.value, error :%s", err.Error())
				    }
				    doc.Value = out
			    }
			
			    err = vdb.mongoDB.SaveDoc(doc)
			    if err != nil {
				    logger.Errorf("Error during Commit(): %s\n", err.Error())
			    }
		    }
	    }
	}

	err = vdb.recordSavepoint(height)
	if err != nil{
		logger.Errorf("Error during recordSavepoint : %s\n", err.Error())
		return err
	}
	
	return nil
}

//savePoint 保存最近记录的信息
func (vdb *VersionedDB) recordSavepoint(height *version.Height) error {
    err := vdb.mongoDB.Delete(savePointNs, savePointKey)
	if err != nil{
		logger.Errorf("Error during delete old savepoint , error : %s", err.Error())
	}
	
	doc := mongodbhelper.MongodbDoc{}
	doc.Version.BlockNum = height.BlockNum
	doc.Version.TxNum = height.TxNum
	doc.Value = nil
	doc.ChaincodeId = savePointNs
	doc.Key = savePointKey
	
	err = vdb.mongoDB.SaveDoc(doc)
	if err != nil{
		logger.Errorf("Error during update savepoint , error : $s", err.Error())
		return err
	}
	
	return err
}

func (vdb *VersionedDB) GetLatestSavePoint() (*version.Height, error) {
	
	doc, err := vdb.mongoDB.GetDoc(savePointNs, savePointKey)
	if err != nil{
		logger.Errorf("Error during get latest save point key, error : %s", err.Error())
		return nil, err
	}
	if doc == nil{
		return nil, nil
	}
	
	return &doc.Version, nil
}

type kvScanner struct {
	namespace string
	result *mgo.Iter
}

func (scanner *kvScanner) Next() (statedb.QueryResult, error) {
    doc := mongodbhelper.MongodbDoc{}
	if !scanner.result.Next(&doc) {
		return nil, nil
	}
	
	blockNum := doc.Version.BlockNum
	txNum := doc.Version.TxNum
	height := version.NewHeight(blockNum, txNum)
	key := doc.Key
	value := doc.Value
	valueContent := []byte{}
	if doc.Value != nil {
		valueContent, _ = json.Marshal(value)
	} else {
		valueContent = doc.Attachments.AttachmentBytes
	}
	
	return &statedb.VersionedKV{
		CompositeKey:   statedb.CompositeKey{Namespace: scanner.namespace, Key: key},
		VersionedValue: statedb.VersionedValue{Value:valueContent , Version: height},
	}, nil
}

func(scanner *kvScanner) Close() {
    err := scanner.result.Close()
	if err != nil{
		logger.Errorf("Error during close the iterator of scanner error : %s", err.Error())
	}
}

func newKVScanner(iter *mgo.Iter, namespace string) *kvScanner {
	return &kvScanner{namespace:namespace, result:iter}
}

func newVersionDB(mgoSession *mgo.Session, dbName string) (*VersionedDB, error) {
    db := mgoSession.DB(dbName)
    conf := mongodbhelper.GetMongoDBConf()
    conf.DBName = dbName
    MongoDB := &mongodbhelper.MongoDB{db, conf}
    return &VersionedDB{MongoDB, dbName}, nil
}

func isJson(value []byte) bool{
	var result interface{}
	err := json.Unmarshal(value, &result)
	return err == nil
}


