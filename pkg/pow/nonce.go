package pow

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"math"
	"math/big"
)

// difficulty of the hashcash algorithm to compute the nonce. The closer to 256,
// the harder to find a nonce.
var difficulty = 16

// ErrNonceNotFound error when a nonce is not found.
var ErrNonceNotFound = errors.New("pow: nonce not found")

// Nonce is the first number that satisfies the hashcat algorithm:
//
// data + nonce < target
type Nonce struct {
	Value   int32  `json:"value"`
	Payload []byte `json:"payload"`
}

// newNonce returns a nonce with its corresponding payload.
func newNonce(data []byte, value int32) *Nonce {
	nonce := Nonce{
		Value:   value,
		Payload: []byte{},
	}
	nonce.computePayload(data)
	return &nonce
}

// computePayload generates the nonce payload by the sum of the nonce value and
// the original data.
func (n *Nonce) computePayload(data []byte) {
	// Convert the data and the nonce to big integer
	dataBigInt := new(big.Int).SetBytes(data)
	nonceBigInt := big.NewInt(int64(n.Value))

	// Create the payload by adding the nonce to the original data
	payloadBigInt := new(big.Int).Add(dataBigInt, nonceBigInt)

	// Apply the sha256 algorithm to the payload
	hash := sha256.Sum256(payloadBigInt.Bytes())
	n.Payload = hash[:]
}

// initTarget will generate the hashcat target based on the difficulty
// parameter.
func initTarget(difficulty int) *big.Int {
	return new(big.Int).Lsh(big.NewInt(1), uint(256-difficulty))
}

// FindNonce will find the nonce as the number that satisfies the hashcash
// algorithm.
func FindNonce(data []byte) (*Nonce, error) {
	// Initialize the target
	target := initTarget(difficulty)

	// Loop until the potencial nonce number (alpha) matches the hashcash
	// condition
	for alpha := int32(0); alpha < math.MaxInt32; alpha++ {
		// create a new test number
		nonce := newNonce(data, alpha)

		// Is the nonce payload smaller than the target number?
		if target.Cmp(new(big.Int).SetBytes(nonce.Payload)) > 0 {
			return nonce, nil
		}
	}

	return nil, ErrNonceNotFound
}

// String prints the nonce in json format.
func (n Nonce) String() string {
	jsonNonce, _ := json.MarshalIndent(n, "", "  ")
	return string(jsonNonce)
}
