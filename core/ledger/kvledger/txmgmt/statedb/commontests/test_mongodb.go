package commontests

import (
    "testing"
    "github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/statedb"
    "github.com/oxchains/fabric/core/ledger/kvledger/txmgmt/version"
    "github.com/oxchains/fabric/common/ledger/testutil"
    "strings"
)

func TestMongoQuery(t *testing.T, dbProvider statedb.VersionedDBProvider) {
    db, err := dbProvider.GetDBHandle("testquery")
    testutil.AssertNoError(t, err, "")
    db.Open()
    defer db.Close()
    batch := statedb.NewUpdateBatch()
    jsonValue1 := "{\"asset_name\": \"marble1\",\"color\": \"blue\",\"size\": 1,\"owner\": \"tom\"}"
    batch.Put("ns1", "key1", []byte(jsonValue1), version.NewHeight(1, 1))
    jsonValue2 := "{\"asset_name\": \"marble2\",\"color\": \"blue\",\"size\": 2,\"owner\": \"jerry\"}"
    batch.Put("ns1", "key2", []byte(jsonValue2), version.NewHeight(1, 2))
    jsonValue3 := "{\"asset_name\": \"marble3\",\"color\": \"blue\",\"size\": 3,\"owner\": \"fred\"}"
    batch.Put("ns1", "key3", []byte(jsonValue3), version.NewHeight(1, 3))
    jsonValue4 := "{\"asset_name\": \"marble4\",\"color\": \"blue\",\"size\": 4,\"owner\": \"martha\"}"
    batch.Put("ns1", "key4", []byte(jsonValue4), version.NewHeight(1, 4))
    jsonValue5 := "{\"asset_name\": \"marble5\",\"color\": \"blue\",\"size\": 5,\"owner\": \"fred\"}"
    batch.Put("ns1", "key5", []byte(jsonValue5), version.NewHeight(1, 5))
    jsonValue6 := "{\"asset_name\": \"marble6\",\"color\": \"blue\",\"size\": 6,\"owner\": \"elaine\"}"
    batch.Put("ns1", "key6", []byte(jsonValue6), version.NewHeight(1, 6))
    jsonValue7 := "{\"asset_name\": \"marble7\",\"color\": \"blue\",\"size\": 7,\"owner\": \"fred\"}"
    batch.Put("ns1", "key7", []byte(jsonValue7), version.NewHeight(1, 7))
    jsonValue8 := "{\"asset_name\": \"marble8\",\"color\": \"blue\",\"size\": 8,\"owner\": \"elaine\"}"
    batch.Put("ns1", "key8", []byte(jsonValue8), version.NewHeight(1, 8))
    jsonValue9 := "{\"asset_name\": \"marble9\",\"color\": \"green\",\"size\": 9,\"owner\": \"fred\"}"
    batch.Put("ns1", "key9", []byte(jsonValue9), version.NewHeight(1, 9))
    jsonValue10 := "{\"asset_name\": \"marble10\",\"color\": \"green\",\"size\": 10,\"owner\": \"mary\"}"
    batch.Put("ns1", "key10", []byte(jsonValue10), version.NewHeight(1, 10))
    jsonValue11 := "{\"asset_name\": \"marble11\",\"color\": \"cyan\",\"size\": 1000007,\"owner\": \"joe\"}"
    batch.Put("ns1", "key11", []byte(jsonValue11), version.NewHeight(1, 11))
    
    //add keys for a separate namespace
    batch.Put("ns2", "key1", []byte(jsonValue1), version.NewHeight(1, 12))
    batch.Put("ns2", "key2", []byte(jsonValue2), version.NewHeight(1, 13))
    batch.Put("ns2", "key3", []byte(jsonValue3), version.NewHeight(1, 14))
    batch.Put("ns2", "key4", []byte(jsonValue4), version.NewHeight(1, 15))
    batch.Put("ns2", "key5", []byte(jsonValue5), version.NewHeight(1, 16))
    batch.Put("ns2", "key6", []byte(jsonValue6), version.NewHeight(1, 17))
    batch.Put("ns2", "key7", []byte(jsonValue7), version.NewHeight(1, 18))
    batch.Put("ns2", "key8", []byte(jsonValue8), version.NewHeight(1, 19))
    batch.Put("ns2", "key9", []byte(jsonValue9), version.NewHeight(1, 20))
    batch.Put("ns2", "key10", []byte(jsonValue10), version.NewHeight(1, 21))
    
    savePoint := version.NewHeight(2, 22)
    db.ApplyUpdates(batch, savePoint)
    
    // query for owner=jerry, use namespace "ns1"
    itr, err := db.ExecuteQuery("ns1", "{\"owner\":\"jerry\"}")
    testutil.AssertNoError(t, err, "")
    
    // verify one jerry result
    queryResult1, err := itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord := queryResult1.(*statedb.VersionedKV)
    stringRecord := string(versionedQueryRecord.Value)
    bFoundRecord := strings.Contains(stringRecord, "jerry")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult2, err := itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult2)
    
    
    // query for owner=jerry, use namespace "ns2"
    itr, err = db.ExecuteQuery("ns2", "{\"owner\":\"jerry\"}")
    testutil.AssertNoError(t, err, "")
    
    // verify one jerry result
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "jerry")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult2, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult2)
    
    // query for owner=jerry, use namespace "ns3"
    itr, err = db.ExecuteQuery("ns3", "{\"owner\":\"jerry\"}")
    testutil.AssertNoError(t, err, "")
    
    // verify results - should be no records
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult1)
    
    // query using bad query string
    itr, err = db.ExecuteQuery("ns1", "this is an invalid query string")
    testutil.AssertError(t, err, "Should have received an error for invalid query string")
    
    // query returns 0 records
    itr, err = db.ExecuteQuery("ns1", "{\"owner\":\"not_a_valid_name\"}")
    testutil.AssertNoError(t, err, "")
    
    // verify no results
    queryResult3, err := itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult3)
    
    // query with complex selector, namespace "ns1"
    itr, err = db.ExecuteQuery("ns1", "{\"$and\":[{\"size\":{\"$gt\": 5}},{\"size\":{\"$lt\":8}},{\"size\":{\"$not\":{\"$eq\":6}}}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify one fred result
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "fred")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult2, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult2)
    
    // query with complex selector, namespace "ns2"
    itr, err = db.ExecuteQuery("ns2", "{\"$and\":[{\"size\":{\"$gt\": 5}},{\"size\":{\"$lt\":8}},{\"size\":{\"$not\":{\"$eq\":6}}}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify one fred result
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "fred")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult2, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult2)
    
    // query with complex selector, namespace "ns3"
    itr, err = db.ExecuteQuery("ns3", "{\"$and\":[{\"size\":{\"$gt\": 5}},{\"size\":{\"$lt\":8}},{\"size\":{\"$not\":{\"$eq\":6}}}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify no more results
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult1)
    
    // query with embedded implicit "AND" and explicit "OR", namespace "ns1"
    itr, err = db.ExecuteQuery("ns1", "{\"color\":\"green\",\"$or\":[{\"owner\":\"fred\"},{\"owner\":\"mary\"}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify one green result
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "green")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify another green result
    queryResult2, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult2)
    versionedQueryRecord = queryResult2.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "green")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult3, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult3)
    
    // query with embedded implicit "AND" and explicit "OR", namespace "ns2"
    itr, err = db.ExecuteQuery("ns2", "{\"color\":\"green\",\"$or\":[{\"owner\":\"fred\"},{\"owner\":\"mary\"}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify one green result
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "green")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify another green result
    queryResult2, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult2)
    versionedQueryRecord = queryResult2.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "green")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult3, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult3)
    
    // query with embedded implicit "AND" and explicit "OR", namespace "ns3"
    itr, err = db.ExecuteQuery("ns3", "{\"color\":\"green\",\"$or\":[{\"owner\":\"fred\"},{\"owner\":\"mary\"}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify no results
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult1)
    
    // query with integer with digit-count equals 7 and response received is also received
    // with same digit-count and there is no float transformation
    itr, err = db.ExecuteQuery("ns1", "{\"$and\":[{\"size\":{\"$eq\": 1000007}}]}")
    testutil.AssertNoError(t, err, "")
    
    // verify one jerry result
    queryResult1, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNotNil(t, queryResult1)
    versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
    stringRecord = string(versionedQueryRecord.Value)
    bFoundRecord = strings.Contains(stringRecord, "joe")
    testutil.AssertEquals(t, bFoundRecord, true)
    bFoundRecord = strings.Contains(stringRecord, "1000007")
    testutil.AssertEquals(t, bFoundRecord, true)
    
    // verify no more results
    queryResult2, err = itr.Next()
    testutil.AssertNoError(t, err, "")
    testutil.AssertNil(t, queryResult2)
}
