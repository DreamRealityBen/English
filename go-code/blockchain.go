package main

func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}
func GenesisBlock() *Block {
	temp := NewBlock("创世块", []byte{})
	return temp
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]

	block := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}
