package mongodbhelper

import (
	"testing"
	"github.com/oxchains/fabric/common/ledger/testutil"
	ledgertestutil "github.com/oxchains/fabric/core/ledger/testutil"
)

func TestGetCouchDBDefinition(t *testing.T) {
	ledgertestutil.SetupCoreYAMLConfig()
	conf := GetMongoDBConf()
	testutil.AssertEquals(t, conf.CollectionName, "test_base")
	testutil.AssertEquals(t, conf.DBName, "stateMongo_test")
}

