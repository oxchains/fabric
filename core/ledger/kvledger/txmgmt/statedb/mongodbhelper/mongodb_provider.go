package mongodbhelper

import (
	"sync"
	"mgo-2"
)

var lastKeyIndicator = byte(0x01)

type Provider struct {
	db *DB
	dbHandles map[string]*DBHandle
	mux sync.Mutex
}

type DBHandle struct {
	DbName string
	Db *DB
}

func NewProvider(conf *Conf) *Provider{
	db := CreateDB(conf)
	db.Open()
	return &Provider{db, make(map[string]*DBHandle), sync.Mutex{}}
}

//TODO 需要修改逻辑,原返回的db有问题.
func (p *Provider) GetDBHandle(dbName string) *DBHandle{
	p.mux.Lock()
	defer p.mux.Unlock()
	dbHandle := p.dbHandles[dbName]
	if dbHandle == nil{
		dialinfo, err := mgo.ParseURL("")
		if err != nil {
			panic(err)
		}
		conf := &Conf{Dialinfo:dialinfo, Database_name:dbName, Collection_name:"test1"}
		db := CreateDB(conf)
		dbHandle = &DBHandle{dbName, db}
		p.dbHandles[dbName] = dbHandle
		db.Open()
		//TODO 添加创建新的数据库方法(后续)
	}
	return dbHandle
}

func (p *Provider) Close(){
	p.db.Close()
}

func (h *DBHandle) GetDoc(ns string, key []byte)(*MongodbDoc, error){
	return h.Db.GetDoc(ns, string(key))
}

func (h *DBHandle) SaveDoc(key []byte, doc MongodbDoc)(error){
	return h.Db.SaveDoc(string(key), doc)
}

func (h *DBHandle) Delete(ns string, key []byte)(error){
	return h.Db.Delete(ns, string(key))
}

func (h *DBHandle) QueryDocuments(query interface{}) (*mgo.Iter, error) {
    return h.Db.QueryDocuments(query)
}

func (h *DBHandle) GetIterator(ns string, startKey []byte, endKey []byte, querylimit int , queryskip int) *mgo.Iter {
	sKey := startKey
	eKey := endKey
	if endKey == nil {
		// replace the last byte 'dbNameKeySep' by 'lastKeyIndicator'
		eKey[len(eKey)-1] = lastKeyIndicator
	}
	logger.Debugf("Getting iterator for range [%#v] - [%#v]", string(sKey), string(eKey))
	return h.Db.GetIterator(ns, string(sKey), string(eKey), querylimit, queryskip)
}


