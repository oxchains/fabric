#!/bin/bash

peer chaincode query -o orderer.example.com:7050 -C mychannel -n mycc -c '{"Args":["query","a"]}'