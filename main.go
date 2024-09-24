package main

import "fmt"

func main() {
	newblockchain := NewBlockchain()
	newblockchain.AddBlock("first transaction")
	newblockchain.AddBlock("second transaction")
	newblockchain.AddBlock("third transaction")
	for _, block := range newblockchain.Blocks {
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)                 // print the hash of the block
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash) // print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.AllData)
	}

}
