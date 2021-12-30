package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// ChainTestSuite stores the parameters to fully test blockchain operations.
type ChainTestSuite struct {
	chain        Chain
	numOfBlocks  uint64
	genesisBlock Block
	suite.Suite
}

// SetupSuite initializes the test suite paremeters.
func (suite *ChainTestSuite) SetupSuite() {
	suite.numOfBlocks = 10
	suite.genesisBlock = Block{
		Hash:     b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
		Data:     "Genesis",
		PrevHash: nil,
		Nonce:    668,
	}
}

// TearDownSuite deletes the backend instance.
func (suite *ChainTestSuite) TearDownSuite() {
	err := suite.chain.Destroy()
	require.NoError(suite.T(), err)
}

// TestAddBlock tests the addition of a new block to the blockchain.
func (suite *ChainTestSuite) TestAddBlock() {
	// Checks if the first block of the chain is the Genesis block.
	firstBlock, err := suite.chain.GetLastBlock()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.genesisBlock, *firstBlock)

	// Add blocks to the blockchain
	for i := uint64(0); i < suite.numOfBlocks; i++ {
		// Add the block to the database
		newBlock, err := suite.chain.AddBlock("this is a testing block")
		require.NoError(suite.T(), err)

		// Retrieve the block from the database
		dbBlock, err := suite.chain.GetBlock(newBlock.Hash)
		require.NoError(suite.T(), err)

		// Check if the returned and stored blocks are the same
		require.Equal(suite.T(), newBlock, dbBlock)
	}
}

// TestErrBlockNotFound checks the error returned when a block is not found.
func (suite *ChainTestSuite) TestErrBlockNotFound() {
	// Find a non-existent block in the blockchain
	_, err := suite.chain.GetBlock([]byte("oblivion"))
	require.Error(suite.T(), err)
	require.Equal(suite.T(), ErrBlockNotFound, err)
}

// TestChainLength checks if the size of the blockchain is correct.
func (suite *ChainTestSuite) TestChainLength() {
	// Check if the the number of elements in the chain matches the length
	chainLength := suite.chain.Length()
	require.Equal(suite.T(), uint64(suite.numOfBlocks+1), chainLength)
}

// TestChainIterator iterates trough all the blocks in the chain to verify that
// the last element and the total length is correct.
func (suite *ChainTestSuite) TestChainIterator() {
	// Iterate trough the whole chain to count the total number of blocks
	var chainLength uint64
	var block *Block
	iterator, err := suite.chain.NewIterator()
	require.NoError(suite.T(), err)
	for iterator.HasNext() {
		block, err = iterator.Next()
		require.NoError(suite.T(), err)
		chainLength++
	}

	// Check if the last returned block is the Genesis block
	require.Equal(suite.T(), suite.genesisBlock, *block)

	// Check if the number of returned elements match the total length of the
	// blockchain
	require.Equal(suite.T(), suite.numOfBlocks+1, chainLength)
}

// TestSliceBlockchain runs the test suite for the slice of blocks
// backend.
func TestSliceBlockchain(t *testing.T) {
	// Initiliaze a new blockchain stored in an slice of blocks
	chain, err := NewSliceChain()
	require.NoError(t, err)

	// Run a new test suite for this blockchain
	sliceChainTestSuite := ChainTestSuite{
		chain: chain,
	}
	suite.Run(t, &sliceChainTestSuite)
}

// TestBadgerBlockchain runs the test suite for the Badger database backend.
func TestBadgerBlockchain(t *testing.T) {
	// Initiliaze a new blockchain stored in a temporal Badger database. It will
	// be removed when the tests are done.
	chain, err := NewBadgerChain("../../test/blockchain/badger")
	require.NoError(t, err)

	// Run a new test suite for this blockchain
	badgerChainTestSuite := ChainTestSuite{
		chain: chain,
	}
	suite.Run(t, &badgerChainTestSuite)
}
