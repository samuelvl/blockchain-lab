package blockchain

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

// b64ToBytes converts an string in base64 into an slice of bytes.
func b64ToBytes(s string) []byte {
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
				Hash:     b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
				Data:     "Genesis",
				PrevHash: []byte{},
				Nonce:    668,
			},
		},
		{
			block: Block{
				Hash:     b64ToBytes("AACNcVOeoodFtxKewQsbSjmMNsSDtRzDNmedx1xaH3Y="),
				Data:     "this is a testing block",
				PrevHash: b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
				Nonce:    157870,
			},
		},
	}

	for _, test := range tests {
		block := NewBlock(test.block.Data, test.block.PrevHash)
		assert.Equal(t, test.block, *block)
	}
}
