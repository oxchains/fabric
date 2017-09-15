package mongodbhelper

import (
	"fmt"
	"sync"
	
	"mgo-2"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/hyperledger/fabric/common/flogging"
	"labix.org/v2/mgo/bson"
)

//TODO 修改存储的格式,单独编写一个struct Doc来存储
//含key,value,chaincodeid,attachments,version
type MongodbDoc struct {
	Key string
	Value interface{}
	ChaincodeId string
	Attachments *Attachment
	Version version.Height
}

type Attachment struct {
	Name            string
	ContentType     string
	Length          uint64
	AttachmentBytes []byte
}

var logger = flogging.MustGetLogger("mongodbhelper")

type dbState int32

const(
	closed dbState = iota
	opened
)

type DB struct {
	conf *Conf
	dbstate dbState
	mux sync.Mutex
	session *mgo.Session
}

type Conf struct {
	Dialinfo *mgo.DialInfo
	Collection_name string
	Database_name string
}

//是否有实际创建数据库
func CreateDB(conf *Conf)*DB  {
	 return &DB{
		 dbstate: closed,
		 conf: conf}
}

func (dbInst *DB) Open(){
	dbInst.mux.Lock()
	defer dbInst.mux.Unlock()
	if(dbInst.dbstate == opened){
		return
	}
	var err error
	dbInst.session ,err = mgo.DialWithInfo(dbInst.conf.Dialinfo)
	if err != nil{
		panic(fmt.Sprintf("Error while trying to activate the session to the configured database ： %s", err))
	}
	dbInst.buildIndex()
	dbInst.dbstate = opened
}

func (dbInst *DB) GetLocalCollection() *mgo.Collection {
    return dbInst.getCollection(dbInst.conf.Database_name, dbInst.conf.Collection_name)
}

func (dbInst *DB) getCollection(dbName, collectionName string) *mgo.Collection {
    return dbInst.session.DB(dbName).C(collectionName)
}

func (dbInst *DB) buildIndex() {
	c := dbInst.GetLocalCollection()
	index := mgo.Index{Key: []string{KEY}, Unique:false, DropDups:false, Background: false}
	//该方法可以实际创建数据库,但是实际创建数据库的方法具体是什么还有待查找.
	err := c.EnsureIndex(index)
	if err != nil{
		logger.Errorf("Error during build index, error : %s", err.Error())
	}
}

func (dbInst *DB) Close(){
	dbInst.mux.Lock()
	defer dbInst.mux.Unlock()
	if dbInst.dbstate == closed{
		return
	}
	dbInst.session.Close()
	dbInst.dbstate = closed
}

func (dbInst *DB) GetDoc(ns, key string) (*MongodbDoc, error){
	logger.Infof("the getdoc key is %s", key)
	collection_inst := dbInst.GetLocalCollection()
	query_result := collection_inst.Find(bson.M{KEY : key, NS : ns})
	num, err := query_result.Count()
	var resultDoc MongodbDoc

	if err != nil{
	    logger.Error(err)
		return nil, err
	}

	if num == 0{
		logger.Debugf("The corresponding value of this key: %s doesn't exist", key)
		return nil, err
	}else if num != 1 {
		logger.Errorf("This key:%s is repeated in the current collection", key)
		return nil, fmt.Errorf("The key %s is repeated", key)
	}else{
		logger.Infof("The corresponding value of this key: %s is extracted successfully", key)
	}
	logger.Infof("the getdoc result len is %d", num)
	query_result.One(&resultDoc)
	
	return &resultDoc, nil
}

//TODO 待修改
func (dbInst *DB) QueryDocuments(query interface{}) (*mgo.Iter, error) {
    collection := dbInst.GetLocalCollection()
	
	queryLimit := 1000
	result := collection.Find(query).Limit(queryLimit).Sort(KEY)
	return result.Iter(), nil
	//需修改
	
}

func (dbInst *DB) SaveDoc(key string, doc MongodbDoc) (error) {
	collecion_inst := dbInst.GetLocalCollection()
	//TODO 是否要区分插入和更新
	collecion_inst.Remove(bson.M{KEY : key, NS : doc.ChaincodeId})
	
	logger.Infof("the doc key is %s", doc.Key)
	
	err := collecion_inst.Insert(&doc)
	if err != nil{
		logger.Errorf("Error in insert the content of key : %s, error : %s",key, err.Error())
		return err
	}
	
	return nil
}

func (dbInst *DB) Delete(ns, key string) error{
	collection_inst := dbInst.GetLocalCollection()
	err := collection_inst.Remove(bson.M{KEY : key, NS : ns})
	if err != nil{
		logger.Errorf("Error %s happened while delete key %s", err.Error(), key)
		return err
	}
	return nil
}


func (dbInst *DB) GetIterator(ns, startkey string, endkey string, querylimit int , queryskip int) *mgo.Iter{
	collection_inst := dbInst.GetLocalCollection()
	query_result := collection_inst.Find(bson.M{KEY : bson.M{"$gte" : startkey, "$lt" : endkey}, NS : ns})
	query_result = query_result.Skip(queryskip)
	query_result = query_result.Limit(querylimit)
	resultRes, _ := query_result.Count()
	logger.Infof("the iteraotr len is %d", resultRes)
	query_result.Sort(KEY)
	return query_result.Iter()

}