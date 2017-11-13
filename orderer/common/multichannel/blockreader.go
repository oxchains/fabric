package multichannel

import (
    cb "github.com/hyperledger/fabric/protos/common"
    "github.com/hyperledger/fabric/orderer/common/ledger"
)

type BlockReader struct {
    blockIndex uint64
    ledgerResources *ledgerResources
    
}

func newBlockReader(ledgerResources *ledgerResources) *BlockReader{
    blockReader := &BlockReader{}
    blockReader.ledgerResources =ledgerResources
    return blockReader
}

func (br *BlockReader) GetBlock(blockIndex uint64) *cb.Block {
    
    block := ledger.GetBlock(br.ledgerResources, blockIndex)
    return block
}