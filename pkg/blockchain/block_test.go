package blockchain

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

// stringToHash converts a hash in base64 into an slice of bytes.
func b64ToHash(s string) []byte {
	hash, _ := base64.StdEncoding.DecodeString(s)
	return hash
}

// TestNewBlock tests the creation of a new block.
func TestNewBlock(t *testing.T) {
	var tests = []struct {
		block Block
	}{
		{
			block: Block{
				Hash: b64ToHash("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
				Data: "Genesis",
			},
		},
		{
			block: Block{
				Hash:     b64ToHash("xL2OQM8Z7a5QloweIkbbBv45sxtX/j4/84h5HmqQxUE="),
				Data:     "this is a testing block",
				PrevHash: b64ToHash("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
			},
		},
	}

	for _, test := range tests {
		block := NewBlock(test.block.Data, test.block.PrevHash)
		assert.Equal(t, test.block, *block)
	}
}
