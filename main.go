package main

import (
	"fmt"

	"github.com/samuelvl/blockchain-lab/pkg/blockchain"
)

func main() {
	chain := blockchain.NewChain()
	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	fmt.Printf("Chain size is %d blocks\n", len(chain.Blocks))

	for index, block := range chain.Blocks {
		fmt.Printf("Block %d is: \n%s\n", index, block)
	}
}
