package mongodbhelper

import (
	"testing"
	"github.com/hyperledger/fabric/common/ledger/testutil"
	ledgertestutil "github.com/hyperledger/fabric/core/ledger/testutil"
)

func TestGetCouchDBDefinition(t *testing.T) {
	ledgertestutil.SetupCoreYAMLConfig()
	conf := GetMongoDBConf()
	testutil.AssertEquals(t, conf.CollectionName, "test_base")
	testutil.AssertEquals(t, conf.DBName, "stateMongo_test")
}

