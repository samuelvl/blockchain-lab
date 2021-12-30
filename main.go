package main

import (
	"fmt"

	"github.com/samuelvl/blockchain-lab/pkg/blockchain"
)

func main() {
	chain, _ := blockchain.NewBadgerChain("/tmp/blockchain")
	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	fmt.Printf("Chain has %d blocks.\n", chain.Length())

	iterator, _ := chain.NewIterator()
	index := uint64(1)
	for iterator.HasNext() {
		block, _ := iterator.Next()
		fmt.Printf("Block %d is: \n%s\n", chain.Length()-index, block)
		index++
	}
}
