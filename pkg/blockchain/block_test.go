package blockchain

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

// b64ToBytes converts an string in base64 into an slice of bytes.
func b64ToBytes(s string) []byte {
	hash, _ := base64.StdEncoding.DecodeString(s)
	return hash
}

// TestComputeHash tests the resulting hash for a single block.
func TestComputeHash(t *testing.T) {
	var tests = []struct {
		block Block
		hash  []byte
	}{
		{
			block: Block{
				Data:     []byte("Genesis"),
				PrevHash: nil,
			},
			hash: b64ToBytes("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
		},
		{
			block: Block{
				Data:     []byte("this is a testing block"),
				PrevHash: b64ToBytes("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
			},
			hash: b64ToBytes("xL2OQM8Z7a5QloweIkbbBv45sxtX/j4/84h5HmqQxUE="),
		},
	}

	for _, test := range tests {
		test.block.ComputeHash()
		require.Equal(t, test.hash, test.block.Hash)
	}
}

// TestNewBlock tests the creation of a new block.
func TestNewBlock(t *testing.T) {
	var tests = []struct {
		block Block
	}{
		{
			block: Block{
				Hash:     b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
				Data:     []byte("Genesis"),
				PrevHash: nil,
				Nonce:    668,
			},
		},
		{
			block: Block{
				Hash:     b64ToBytes("AACNcVOeoodFtxKewQsbSjmMNsSDtRzDNmedx1xaH3Y="),
				Data:     []byte("this is a testing block"),
				PrevHash: b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
				Nonce:    157870,
			},
		},
	}

	for _, test := range tests {
		block := NewBlock(test.block.Data, test.block.PrevHash)
		require.Equal(t, test.block, *block)
	}
}

// TestBlockSerialization test the serialization and deserialization of a block.
func TestBlockSerialization(t *testing.T) {
	var tests = []struct {
		block Block
	}{
		{
			block: Block{
				Hash:     b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
				Data:     []byte("Genesis"),
				PrevHash: nil,
				Nonce:    668,
			},
		},
	}

	for _, test := range tests {
		// Serialize the block to test
		serializedBlock, err := test.block.Serialize()
		require.NoError(t, err)

		// Deserialize the serialized block
		deserializedBlock := Block{}
		deserializedBlock.Deserialize(serializedBlock)

		// Compare the deserialize with the original one
		require.Equal(t, test.block, deserializedBlock)
	}
}
