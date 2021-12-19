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
					Hash:     b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
					Data:     "Genesis",
					PrevHash: []byte{},
					Nonce:    668,
				},
				{
					Hash:     b64ToBytes("AACNcVOeoodFtxKewQsbSjmMNsSDtRzDNmedx1xaH3Y="),
					Data:     "this is a testing block",
					PrevHash: b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
					Nonce:    157870,
				},
			},
		},
	}

	chain := NewChain()
	assert.Equal(t, test.chain.Blocks[0], chain.Blocks[0])

	chain.AddBlock(test.chain.Blocks[1].Data)
	assert.Equal(t, test.chain.Blocks[1], chain.Blocks[1])
}
