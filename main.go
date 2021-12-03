package main

import (
	"fmt"

	"github.com/samuelvl/blockchain-lab/pkg/blockchain"
)

func main() {
	block := blockchain.FirstBlock()
	fmt.Printf("Block: \n%s", block)
}
