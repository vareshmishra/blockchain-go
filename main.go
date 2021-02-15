package main

import (
	"fmt"

	"github.com/vareshmishra/blockchain-go/blockchain"
)

func main() {
	fmt.Println("Testing Blockchain")

	chain := blockchain.InitBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, blk := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", blk.PrevHash)
		fmt.Printf("Data in Block: %s\n", blk.Data)
		fmt.Printf("Hash: %x\n", blk.Hash)
		fmt.Println("================")

	}
}
