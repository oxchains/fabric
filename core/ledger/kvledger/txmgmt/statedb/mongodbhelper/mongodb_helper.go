package mongodbhelper

import (
	"fmt"
	
	"mgo-2"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/hyperledger/fabric/common/flogging"
	"labix.org/v2/mgo/bson"
)

//保存在mongodb中的文档格式
type MongodbDoc struct {
	Key string
	Value interface{}
	ChaincodeId string
	Attachments *Attachment
	Version version.Height
}

//当value不是json时,则保存在attachment
type Attachment struct {
	Name            string
	ContentType     string
	Length          uint64
	AttachmentBytes []byte
}

var logger = flogging.MustGetLogger("mongodbhelper")

type MongoDB struct {
	Db *mgo.Database
	Conf *MongoDBConf
}

func (mongoDB *MongoDB) GetDefaultCollection() *mgo.Collection {
    return mongoDB.getCollection(mongoDB.Conf.CollectionName)
}

func (mongoDB *MongoDB) getCollection(collectionName string) *mgo.Collection {
    return mongoDB.Db.C(collectionName)
}

//建立索引,暂未使用.
func (mongoDB *MongoDB) buildIndex() {
	c := mongoDB.GetDefaultCollection()
	index := mgo.Index{Key: []string{KEY}, Unique:false, DropDups:false, Background: false}
	err := c.EnsureIndex(index)
	if err != nil{
		logger.Errorf("Error during build index, error : %s", err.Error())
	}
}

func (mongoDb *MongoDB) Open() {

}

func (mongoDB *MongoDB) Close() {

}

func (mongoDB *MongoDB) GetDoc(ns, key string) (*MongodbDoc, error) {
	collection := mongoDB.GetDefaultCollection()
	queryResult := collection.Find(bson.M{KEY : key, NS : ns})
	num, err := queryResult.Count()
	if err != nil{
		logger.Error(err)
		return nil, err
	}
	
	var resultDoc MongodbDoc
	if num == 0 {
		logger.Debugf("The corresponding value of this key: %s doesn't exist", key)
		return nil, nil
	} else if num != 1 {
		logger.Errorf("This key:%s is repeated in the current collection", key)
		return nil, fmt.Errorf("The key %s is repeated", key)
	}
	queryResult.One(&resultDoc)
	return &resultDoc, nil
}

func (mongoDB *MongoDB) QueryDocuments(query interface{}) (*mgo.Iter, error) {
    collection := mongoDB.GetDefaultCollection()
	queryLimit := mongoDB.Conf.QueryLimit
	result := collection.Find(query).Limit(queryLimit).Sort(KEY)
	return result.Iter(), nil
}

func (mongoDB *MongoDB) SaveDoc(doc MongodbDoc) (error) {
	collection := mongoDB.GetDefaultCollection()
	//TODO 是否要区分插入和更新
	collection.Remove(bson.M{KEY : doc.Key, NS : doc.ChaincodeId})
	err := collection.Insert(&doc)
	if err != nil{
		logger.Errorf("Error in insert the content of key : %s, error : %s",doc.Key, err.Error())
		return err
	}
	
	return nil
}

func (mongoDB *MongoDB) Delete(ns, key string) error {
	collection := mongoDB.GetDefaultCollection()
	err := collection.Remove(bson.M{KEY : key, NS : ns})
	if err != nil{
		logger.Errorf("Error %s happened while delete key %s", err.Error(), key)
		return err
	}
	
	return nil
}


func (mongoDB *MongoDB) GetIterator(ns, startkey string, endkey string, querylimit int , queryskip int) *mgo.Iter {
	collection := mongoDB.GetDefaultCollection()
	var queryResult *mgo.Query
	if endkey == "" {
		queryResult = collection.Find(bson.M{KEY : bson.M{"$gte" : startkey}, NS : ns})
	} else {
		queryResult = collection.Find(bson.M{KEY : bson.M{"$gte" : startkey, "$lt" : endkey}, NS : ns})
	}
	
	queryResult = queryResult.Skip(queryskip)
	queryResult = queryResult.Limit(querylimit)
	queryResult.Sort(KEY)
	return queryResult.Iter()
}