package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAddBlock tests the addition of a new block to the chain.
func TestAddBlock(t *testing.T) {
	var test = struct {
		chain Chain
	}{
		chain: Chain{
			Blocks: []*Block{
				{
					Hash:     b64ToHash("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
					Data:     "Genesis",
					PrevHash: []byte{},
				},
				{
					Hash:     b64ToHash("xL2OQM8Z7a5QloweIkbbBv45sxtX/j4/84h5HmqQxUE="),
					Data:     "this is a testing block",
					PrevHash: b64ToHash("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
				},
			},
		},
	}

	chain := NewChain()
	assert.Equal(t, test.chain.Blocks[0], chain.Blocks[0])

	chain.AddBlock(test.chain.Blocks[1].Data)
	assert.Equal(t, test.chain.Blocks[1], chain.Blocks[1])
}
