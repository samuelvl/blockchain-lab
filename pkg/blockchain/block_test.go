package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestComputeHash tests the resulting hash for a single block.
func TestComputeHash(t *testing.T) {
	var tests = []struct {
		block Block
		hash  string
	}{
		{
			block: Block{
				Data:     []byte("Genesis"),
				PrevHash: "",
			},
			hash: "81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5",
		},
		{
			block: Block{
				Data:     []byte("this is a testing block"),
				PrevHash: "81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5",
			},
			hash: "5546e8962b45ef7d89ec93e54162bca55129914c1766d2fb0c74492f1f9ec776",
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
				Data:     []byte("Genesis"),
				Hash:     "0000f5adf42baf5174fc801e930ab3d020b5d00218657e66df8f23419da9c3c1",
				PrevHash: "",
				Nonce:    205317,
			},
		},
		{
			block: Block{
				Data:     []byte("this is a testing block"),
				Hash:     "00005bfbe5cfb03a5d9e729d884700ebaa1cf825d9be9790636d991d03b51e18",
				PrevHash: "0000f5adf42baf5174fc801e930ab3d020b5d00218657e66df8f23419da9c3c1",
				Nonce:    107902,
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
				Data:     []byte("Genesis"),
				Hash:     "81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5",
				PrevHash: "",
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
