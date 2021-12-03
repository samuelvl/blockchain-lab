package blockchain

import "sync"

// Chain represents an slice of blocks.
type Chain struct {
	Blocks []*Block
	sync.Mutex
}

// NewChain initializes the chain with a first single "genesis" block.
func NewChain() *Chain {
	chain := Chain{
		Blocks: []*Block{FirstBlock()},
	}
	return &chain
}

// AddBlock adds a new block to the chain from the input data.
func (c *Chain) AddBlock(data string) {
	// Avoid race conditions while adding new blocks
	c.Lock()
	defer c.Unlock()

	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}
