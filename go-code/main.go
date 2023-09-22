package main

import (
	"fmt"
)

func main() {

	bc := NewBlockChain()
	bc.AddBlock("Ben transfer one bitcoin to Allen")

	for i, block := range bc.blocks {
		fmt.Printf("================BlockHigh: %d ====================\n", i)
		fmt.Printf("prevHash: %x \n", block.PrevHash)
		fmt.Printf("Hash: %x \n", block.Hash)
		fmt.Printf("Data: %s \n", block.Data)
	}

	// block := NewBlock("Blb transfer one bitcoin  to Alen", []byte{})

}
