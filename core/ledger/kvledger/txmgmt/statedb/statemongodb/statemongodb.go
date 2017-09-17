package statemongodb

import (
    "bytes"
	"mgo-2"
	"fmt"
	"unicode/utf8"
	"encoding/json"
	
	"github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/statedb/mongodbhelper"
	"github.com/oxchains/fabric/common/flogging"
	"github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/oxchains/fabric/core/ledger/ledgerconfig"
	ledgertestutil "github.com/oxchains/fabric/core/ledger/testutil"
)


var logger = flogging.MustGetLogger("statemongodb")

var compositeKeySep = []byte{0x00}
var lastKeyIndicator = byte(0x01)
var savePointKey = "statedb_savepoint"
var savePointNs = "savepoint"

var queryskip = 0


type VersionedDBProvider struct {
	dbProvider *mongodbhelper.Provider
}

type VersionedDB struct {
	db *mongodbhelper.DBHandle
	dbName string
}

func NewVersionedDBProvider() *VersionedDBProvider{
	ledgertestutil.SetupCoreYAMLConfig()
	conf := GetMongoDBConf()
	//fmt.Println(conf.Dialinfo.Addrs)
	provider := mongodbhelper.NewProvider(conf)
	return &VersionedDBProvider{provider}

}


func (provider *VersionedDBProvider) GetDBHandle(dbName string) (statedb.VersionedDB, error){
	return &VersionedDB{provider.dbProvider.GetDBHandle(dbName), dbName}, nil
}
func (provider *VersionedDBProvider) Close(){
	provider.dbProvider.Close()
}

func (vdb *VersionedDB) Close(){
	vdb.db.Db.Close()
}

func (vdb *VersionedDB)Open() error{
        vdb.db.Db.Open()
	return nil
}

func (vdb *VersionedDB) GetState(namespace string, key string)(*statedb.VersionedValue, error){
	logger.Infof("GetState(). ns=%s, key = %s", namespace, key)
	compositeKey := constructCompositeKey(namespace, key)
	logger.Infof("getstate(). the compositekey = [%#v]", compositeKey)
	doc , err := vdb.db.GetDoc(namespace, compositeKey)
	if err != nil{
		return nil, err
	}
	if doc == nil{
		return nil, nil
	}
	
	versioned_v := statedb.VersionedValue{Value: nil, Version:& doc.Version,}
	if doc.Value == nil && doc.Attachments != nil {
		versioned_v.Value = doc.Attachments.AttachmentBytes
	}else{
		tempbyte , _ := json.Marshal(doc.Value)
		versioned_v.Value = []byte(string(tempbyte))
	}
	return &versioned_v, nil
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
	compositeStartKey := constructCompositeKey(namespace, startKey)
	compositeEndKey := constructCompositeKey(namespace, endKey)
	querylimit := ledgerconfig.GetQueryLimit()
	if endKey == "" {
		compositeEndKey[len(compositeEndKey)-1] = lastKeyIndicator
	}
	dbItr := vdb.db.GetIterator(namespace, compositeStartKey, compositeEndKey, querylimit, queryskip)
	return newKVScanner(dbItr, namespace), nil
}

func (vdb *VersionedDB) ExecuteQuery(namespace ,query string)(statedb.ResultsIterator, error){
	
	//Get the querylimit from core.yaml
	queryByte := []byte(query)
	if !isJson(queryByte) {
	    return nil, fmt.Errorf("the query is not a json : %s", query)
	}
	//TODO 待修改
	queryBson, err := mongodbhelper.GetQueryBson(namespace, query)
	if err != nil {
	    return nil, err
	}
	queryJson, _ := json.Marshal(queryBson)
	logger.Infof("the queryBson is %s", string(queryJson))
	result, err := vdb.db.QueryDocuments(queryBson)
	if err != nil {
	    return nil, err
	}
	logger.Infof("the queryresult len is %d")
	return newKVScanner(result, namespace), nil
}

//TODO 修改此方法
func(vdb VersionedDB) ApplyUpdates(batch *statedb.UpdateBatch, height *version.Height) error{
    namespaces := batch.GetUpdatedNamespaces()
	var out interface{}
	var err error
	for _, ns := range namespaces{
	    updates := batch.GetUpdates(ns)
	    for k, vv := range updates{
		    compositeKey := constructCompositeKey(ns, k)
		    logger.Infof("Channel [%s]: Applying key=[%#v]", vdb.dbName, compositeKey)
		    if vv.Value == nil{
			    vdb.db.Delete(ns, compositeKey)
		    }else{
		        doc := mongodbhelper.MongodbDoc {
		    		Key: string(compositeKey),
		    		Version: *vv.Version,
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
				    //fmt.Println(out, "in apply")
				    if err != nil{
					    logger.Errorf("Error rises while unmarshal the vv.value, error :%s", err.Error())
				    }
				    doc.Value = out
			    }
			  
			    err = vdb.db.SaveDoc(compositeKey, doc)
			    if err != nil {
				    logger.Errorf("Error during Commit(): %s\n", err.Error())
			    }
		    }
	  }
	}
	//该步骤不需要命名空间吗
	err = vdb.recordSavepoint(height)
	if err != nil{
		logger.Errorf("Error during recordSavepoint : %s\n", err.Error())
		return err
	}
	return nil
}

//savePoint 保存最近记录的信息
func (vdb *VersionedDB) recordSavepoint(height *version.Height) error{
    err := vdb.db.Delete(savePointNs, []byte(savePointKey))
	if err != nil{
		logger.Errorf("Error during delete old savepoint , error : %s", err.Error())
	}
	
	doc := mongodbhelper.MongodbDoc{}
	doc.Version.BlockNum = height.BlockNum
	doc.Version.TxNum = height.TxNum
	doc.Value = nil
	doc.ChaincodeId = savePointNs
	doc.Key = savePointKey
	
	err = vdb.db.SaveDoc([]byte(savePointKey), doc)
	if err != nil{
		logger.Errorf("Error during update savepoint , error : $s", err.Error())
		return err
	}
	
	return err
}

func (vdb *VersionedDB) GetLatestSavePoint() (*version.Height, error){
	doc, err := vdb.db.GetDoc(savePointNs, []byte(savePointKey))
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
	if !scanner.result.Next(&doc){
		return nil, nil
	}
	blockNum := doc.Version.BlockNum
	txNum := doc.Version.TxNum
	height := version.NewHeight(blockNum, txNum)
	key := doc.Key
	_, key = splitCompositeKey([]byte(key))
	value := doc.Value
	value_content := []byte{}
	if doc.Value != nil{
	    value_content, _ = json.Marshal(value)
	} else {
	    value_content = doc.Attachments.AttachmentBytes
	}
	
	//测试输出doc
	docJson, _ := json.Marshal(doc)
	logger.Infof("the docjson is %s", string(docJson))
	//测试结束
	
	return &statedb.VersionedKV{
		CompositeKey:   statedb.CompositeKey{Namespace: scanner.namespace, Key: key},
		VersionedValue: statedb.VersionedValue{Value:value_content , Version: height},
	}, nil
}

func(scanner *kvScanner) Close(){
    err := scanner.result.Close()
	if err != nil{
		logger.Errorf("Error during close the iterator of scanner error : %s", err.Error())
	}

}

func newKVScanner(iter *mgo.Iter, namespace string) *kvScanner{
	return &kvScanner{namespace:namespace, result:iter}
}

func constructCompositeKey(ns string, key string) []byte {
	return append(append([]byte(ns), compositeKeySep...), []byte(key)...)
}

func splitCompositeKey(compositeKey []byte) (string, string) {
	split := bytes.SplitN(compositeKey, compositeKeySep, 2)
	return string(split[0]), string(split[1])
}

func isJson(value []byte) bool{
	var result interface{}
	err := json.Unmarshal(value, &result)
	return err == nil
}


