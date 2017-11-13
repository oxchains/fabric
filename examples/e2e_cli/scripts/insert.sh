#!/bin/bash

# 测试orderer脚本

orderer_0=orderer0.example.com:7050
orderer_1=orderer1.example.com:7050
orderer_2=orderer2.example.com:7050
orderer_3=orderer3.example.com:7050
channel_file_path=./channel-artifacts/channel.tx
channel_name=mychannel

peer channel create -f $channel_file_path -o $orderer_0 -c $channel_name
sleep 1
peer channel join -b $channel_name.block
sleep 1

peer chaincode install -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02 -v 1.0 -n mycc

sleep 1

peer chaincode instantiate -n mycc -v 1.0 -o orderer0.example.com:7050 -C mychannel -c '{"Args":["","a","100","b","100"]}'

sleep 5

query() {
    orderer=$1
    arg=$2
    shouldbe=$3
    peer chaincode query -n mycc -o $orderer -C $channel_name -c '{"Args":["query","'$arg'"]}'

    res=$1
    "The result should be $shouldbe, and the really is $res"
    sleep 1
}

invoke() {
    orderer=$1
    arg=$2
    peer chaincode invoke -n mycc -o $orderer -C $channel_name -c '{"Args":["invoke","a","b","'$arg'"]}'
}

query $orderer_0 a 100

invoke $orderer_0 10

invoke $orderer_0 20

invoke $orderer_0 30

invoke $orderer_0 40

invoke $orderer_0 50

invoke $orderer_0 60

#query $orderer_0 a 80

#query $orderer_0 b 120
