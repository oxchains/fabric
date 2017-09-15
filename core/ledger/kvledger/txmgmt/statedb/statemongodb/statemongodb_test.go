package statemongodb

import (
	"testing"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/commontests"
	"github.com/hyperledger/fabric/common/ledger/testutil"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
)

func TestBasicRW(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testbasicrw")
	commontests.TestBasicRW(t, env.DBProvider)
}

func TestGetStateMultipleKeys(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testgetmultiplekeys")
	commontests.TestGetStateMultipleKeys(t, env.DBProvider)
}

func TestMultiDBBasicRW(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testmultidbbasicrw")
	defer env.Cleanup("testmultidbbasicrw2")
	commontests.TestMultiDBBasicRW(t, env.DBProvider)
}

func TestDeletes(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testdeletes")
	commontests.TestDeletes(t, env.DBProvider)
}

func TestEncodeDecodeValueAndVersion(t *testing.T) {
	testValueAndVersionEncodeing(t, []byte("value1"), version.NewHeight(1, 2))
	testValueAndVersionEncodeing(t, []byte{}, version.NewHeight(50, 50))
}

func testValueAndVersionEncodeing(t *testing.T, value []byte, version *version.Height) {
	encodedValue := statedb.EncodeValue(value, version)
	val, ver := statedb.DecodeValue(encodedValue)
	testutil.AssertEquals(t, val, value)
	testutil.AssertEquals(t, ver, version)
}

func TestIterator(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testiterator")
	commontests.TestIterator(t, env.DBProvider)
	
}

func TestJsonQuery(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testquery")
	commontests.TestMongoQuery(t, env.DBProvider)
}
