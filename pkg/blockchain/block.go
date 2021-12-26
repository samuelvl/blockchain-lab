package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"

	"github.com/samuelvl/blockchain-lab/pkg/pow"
)

// Difficulty of the hashcash algorithm to compute the nonce. The closer to 256,
// the harder to find a nonce.
const Difficulty uint = 16

// Block represents the simplest element of the chain. It contains an string,
// its corresponding hash and the hash from the previous block.
// The previous hash will be empty if it is the first block of the chain.
type Block struct {
	Hash     []byte `json:"hash"`
	Data     string `json:"data"`
	PrevHash []byte `json:"prevHash"`
	Nonce    int32  `json:"nonce"`
}

// NewBlock returns a block with its corresponding hash.
func NewBlock(data string, prevHash []byte) *Block {
	block := Block{
		Hash:     []byte{},
		Data:     data,
		PrevHash: prevHash,
		Nonce:    0,
	}
	block.ComputeHash()
	block.Mine()
	return &block
}

// FirstBlock returns the first block of the chain from the "Genesis" string.
func FirstBlock() *Block {
	return NewBlock("Genesis", []byte{})
}

// ComputeHash computes block's hash using the sha256 algorithm:
// https://datatracker.ietf.org/doc/html/rfc6234
func (b *Block) ComputeHash() {
	// The payload is the concatenation of the block's data and the previous
	// hash, this is join[data, padding, prevHash]. No padding is added between
	// the data and the previous hash.
	padding := []byte{}
	payload := bytes.Join([][]byte{[]byte(b.Data), b.PrevHash}, padding)

	// Set the hash value as the sha256 of the payload
	hash := sha256.Sum256(payload)
	b.Hash = hash[:]
}

// Mine will recompute the block's hash using the Proof of Work "hashcat"
// algorithm.
func (b *Block) Mine() error {
	nonce, err := pow.FindNonce(b.Hash, Difficulty)
	if err != nil {
		return err
	}
	b.Hash = nonce.Payload
	b.Nonce = nonce.Value

	return nil
}

// String prints the block in json format.
func (b Block) String() string {
	jsonBlock, _ := json.MarshalIndent(b, "", "  ")
	return string(jsonBlock)
}
