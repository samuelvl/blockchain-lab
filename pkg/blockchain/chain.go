package blockchain

import (
	"bytes"
	"errors"
	"os"
	"sync"

	badger "github.com/dgraph-io/badger/v3"
)

// ErrBlockNotFound error when a block is not found.
var ErrBlockNotFound = errors.New("blockchain: block not found")

// Chain is the interface to be implemented by a blockchain backend.
type Chain interface {
	AddBlock(data []byte) (*Block, error)
	GetBlock(hash string) (*Block, error)
	GetLastBlock() (*Block, error)
	Destroy() error
	Length() uint64
	NewIterator() (*ChainIterator, error)
}

// ChainIterator can be used to iterate through the blockchain using the Next()
// method.
type ChainIterator struct {
	currentHash string
	chain       Chain
}

// SliceChain will use an slice of blocks as the blockchain backend.
type SliceChain struct {
	Blocks []*Block
	sync.Mutex
}

// NewSliceChain initializes a blockchain to store blocks in an slice of blocks.
// It will add the Genesis block as the first block of the chain.
func NewSliceChain() (*SliceChain, error) {
	chain := SliceChain{
		Blocks: []*Block{FirstBlock()},
	}
	return &chain, nil
}

// AddBlock adds a new block to the chain from the input data.
func (chain *SliceChain) AddBlock(data []byte) (*Block, error) {
	// Avoid race conditions while adding new blocks
	chain.Lock()
	defer chain.Unlock()

	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)

	return newBlock, nil
}

// GetBlock finds and returns a block from its hash. If block is not found,
// ErrBlockNotFound is returned.
func (chain *SliceChain) GetBlock(hash string) (*Block, error) {
	// Avoid race conditions while iterating blocks
	chain.Lock()
	defer chain.Unlock()

	// Iterate the slice of blocks and return the block matching the hash
	for _, block := range chain.Blocks {
		if hash == block.Hash {
			return block, nil
		}
	}

	return nil, ErrBlockNotFound
}

// GetLastBlock returns the last block of the chain.
func (chain *SliceChain) GetLastBlock() (*Block, error) {
	// Avoid race conditions while iterating blocks
	chain.Lock()
	defer chain.Unlock()

	// The last block is in the last position of the slice
	lastBlock := chain.Blocks[len(chain.Blocks)-1]

	return lastBlock, nil
}

// Destroy removes all the blocks from the chain.
func (chain *SliceChain) Destroy() error {
	// Avoid race conditions while iterating blocks
	chain.Lock()
	defer chain.Unlock()

	// The last block is in the last position of the slice
	chain.Blocks = []*Block{}

	return nil
}

// Length returns the total size of the blockchain.
func (chain *SliceChain) Length() uint64 {
	// Avoid race conditions while iterating blocks
	chain.Lock()
	defer chain.Unlock()

	// Count the number of elements
	size := uint64(len(chain.Blocks))

	return size
}

// NewIterator initializes the blockchain iterator from the last block.
func (chain *SliceChain) NewIterator() (*ChainIterator, error) {
	lastBlock, err := chain.GetLastBlock()
	if err != nil {
		return nil, err
	}
	iterator := ChainIterator{
		currentHash: lastBlock.Hash,
		chain:       chain,
	}
	return &iterator, nil
}

// BadgerChain will use a Badger database as the blockchain backend. Badger
// documentation: https://dgraph.io/docs/badger
type BadgerChain struct {
	db           *badger.DB
	lastBlockKey []byte
}

// NewBadgerChain initializes a blockchain to store blocks in a Badger database.
// It will add the Genesis block as the first block of the chain.
func NewBadgerChain(dir string) (*BadgerChain, error) {
	// Create a new badger instance
	config := badger.DefaultOptions(dir)
	config.Logger = nil
	database, err := badger.Open(config)
	if err != nil {
		return nil, err
	}

	// Configure the Badger database as the blockchain backend
	chain := BadgerChain{
		db:           database,
		lastBlockKey: []byte("lastBlock"),
	}

	// If the database is not initialized yet, create the Genesis block as the
	// first block
	txn := chain.db.NewTransaction(true)
	defer txn.Discard()

	_, err = txn.Get(chain.lastBlockKey)
	if err == badger.ErrKeyNotFound {
		// Create the Genesis block
		firstBlock := FirstBlock()
		firstBlockBytes, err := firstBlock.Serialize()
		if err != nil {
			return nil, err
		}

		// Add the Genesis block to the chain
		err = txn.SetEntry(
			badger.NewEntry([]byte(firstBlock.Hash), firstBlockBytes))
		if err != nil {
			return nil, err
		}

		// Update the last block key with the Genesis block
		err = txn.SetEntry(
			badger.NewEntry(chain.lastBlockKey, firstBlockBytes))
		if err != nil {
			return nil, err
		}

		// Commit the transaction and check for error
		err = txn.Commit()
		if err != nil {
			return nil, err
		}
	}

	return &chain, nil
}

// AddBlock adds a new block to the chain from the input data.
func (chain *BadgerChain) AddBlock(data []byte) (*Block, error) {
	// Create a new read-write badger transaction
	txn := chain.db.NewTransaction(true)
	defer txn.Discard()

	// Get the previous block from the last block key
	var prevBlock Block
	prevBlockBytes, err := txn.Get(chain.lastBlockKey)
	if err != nil {
		return nil, err
	}

	prevBlockRaw, err := prevBlockBytes.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	err = prevBlock.Deserialize(prevBlockRaw)
	if err != nil {
		return nil, err
	}

	// Create the new block from the previous block hash
	var block *Block
	block = NewBlock(data, prevBlock.Hash)
	blockBytes, err := block.Serialize()
	if err != nil {
		return nil, err
	}

	// Add the new block to the database
	err = txn.SetEntry(badger.NewEntry([]byte(block.Hash), blockBytes))
	if err != nil {
		return nil, err
	}

	// Update the last block key
	err = txn.SetEntry(badger.NewEntry(chain.lastBlockKey, blockBytes))
	if err != nil {
		return nil, err
	}

	// Commit the transaction and check for error
	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return block, nil
}

// GetBlock finds and returns a block from its hash. If block is not found,
// ErrBlockNotFound is returned.
func (chain *BadgerChain) GetBlock(hash string) (*Block, error) {
	// Create a new read-only badger transaction
	txn := chain.db.NewTransaction(false)
	defer txn.Discard()

	// Find the block in the database
	var block Block
	blockBytes, err := txn.Get([]byte(hash))
	if err != nil {
		return nil, ErrBlockNotFound
	}

	blockRaw, err := blockBytes.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	err = block.Deserialize(blockRaw)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

// GetLastBlock returns the last block of the chain.
func (chain *BadgerChain) GetLastBlock() (*Block, error) {
	lastBlock, err := chain.GetBlock(string(chain.lastBlockKey))
	if err != nil {
		return nil, err
	}
	return lastBlock, nil
}

// Destroy removes all the blocks from the chain.
func (chain *BadgerChain) Destroy() error {
	err := chain.db.DropAll()
	if err != nil {
		return err
	}
	err = chain.db.Close()
	if err != nil {
		return err
	}
	err = os.RemoveAll(chain.db.Opts().Dir)
	return err
}

// NewIterator initializes the blockchain iterator from the last block.
func (chain *BadgerChain) NewIterator() (*ChainIterator, error) {
	lastBlock, err := chain.GetLastBlock()
	if err != nil {
		return nil, err
	}
	iterator := ChainIterator{
		currentHash: lastBlock.Hash,
		chain:       chain,
	}
	return &iterator, nil
}

// Length returns the total size of the blockchain.
func (chain *BadgerChain) Length() uint64 {
	// Create a new read-only badger transaction
	txn := chain.db.NewTransaction(false)
	defer txn.Discard()

	// Use the Badger iterator to count the number of keys in the database
	var size uint64
	opts := badger.DefaultIteratorOptions
	opts.PrefetchValues = false
	iterator := txn.NewIterator(opts)
	defer iterator.Close()
	for iterator.Rewind(); iterator.Valid(); iterator.Next() {
		// Do not count the last block key
		if bytes.Compare(iterator.Item().Key(), chain.lastBlockKey) != 0 {
			size++
		}
	}

	return size
}

// Next returns the next block in the blockchain until the Genesis block is
// reached.
func (iterator *ChainIterator) Next() (*Block, error) {
	// Get the next element in the blockchain
	nextBlock, err := iterator.chain.GetBlock(iterator.currentHash)
	if err != nil {
		return nil, err
	}

	// Update the iterator current hash pointer
	iterator.currentHash = nextBlock.PrevHash

	return nextBlock, nil
}

// HasNext chechks if the blockchain has remanining blocks.
func (iterator *ChainIterator) HasNext() bool {
	return iterator.currentHash != ""
}
