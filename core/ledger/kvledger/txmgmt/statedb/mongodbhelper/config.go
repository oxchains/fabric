package mongodbhelper

import (
	"github.com/spf13/viper"
	"time"
)

type MongoDBConf struct {
	Url string
	UserName string
	Password string
	DBName string
	CollectionName string
	QueryLimit int
	RequestTimeout time.Duration
}

func GetMongoDBConf() *MongoDBConf {
	url := viper.GetString("ledger.state.MongoDBConfig.url")
	userName := viper.GetString("ledger.state.MongoDBConfig.username")
	password := viper.GetString("ledger.state.MongoDBConfig.password")
	collectionName := viper.GetString("ledger.state.mongoDBConfig.collectionName")
	timeout := viper.GetDuration("ledger.state.mongoDBConfig.requestTimeout")
	
	if collectionName == "" {
		collectionName = "test"
	}
	dbName := viper.GetString("ledger.state.mongoDBConfig.databaseName")
	if dbName == "" {
		dbName = "mongotest"
	}
	queryLimit := viper.GetInt("ledger.state.mongoDBConfig.queryLimit")
	if queryLimit <= 0 {
		queryLimit = 1000
	}
	//timeout=0 代码一直等待服务器响应
	if timeout <= 0 {
		timeout, _ = time.ParseDuration("35s")
	}
	
	return &MongoDBConf{
		Url: url,
		UserName: userName,
		Password: password,
		DBName: dbName,
		CollectionName: collectionName,
		QueryLimit: queryLimit,
		RequestTimeout: timeout,
	}
}