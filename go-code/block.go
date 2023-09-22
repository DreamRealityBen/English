package main

import (
	"crypto/sha256"
)

type Block struct {
	Version uint64,
	PrevHash []byte
	MerkelRoot []byte,
	TimeStamp uint64,
	Diffculty uint64,
	Nonce uint64,

	Hash     []byte
	Data     []byte
}

type BlockChain struct {
	blocks []*Block
}

// create block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}

func (block *Block) SetHash() {
	blockInfo := append(block.PrevHash, block.Data...)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
