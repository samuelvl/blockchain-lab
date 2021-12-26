package pow

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// b64ToBytes converts an string in base64 into an slice of bytes.
func b64ToBytes(s string) []byte {
	data, _ := base64.StdEncoding.DecodeString(s)
	return data
}

// TestInitTarget tests the initialization of a target.
func TestInitTarget(t *testing.T) {
	var tests = []struct {
		difficulty int
		target     string
	}{
		{
			difficulty: 16,
			target:     "0x1000000000000000000000000000000000000000000000000000000000000",
		},
		{
			difficulty: 64,
			target:     "0x1000000000000000000000000000000000000000000000000",
		},
		{
			difficulty: 128,
			target:     "0x100000000000000000000000000000000",
		},
		{
			difficulty: 250,
			target:     "0x40", // 0x1000000 binary
		},
	}

	for _, test := range tests {
		target := initTarget(test.difficulty)
		targetHex := fmt.Sprintf("0x%x", target)
		assert.Equal(t, test.target, targetHex)
	}
}

// TestFindNonce tests the match of a nonce.
func TestFindNonce(t *testing.T) {
	var tests = []struct {
		data  []byte
		nonce Nonce
	}{
		{
			data: b64ToBytes("gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU="),
			nonce: Nonce{
				Value:   668,
				Payload: b64ToBytes("AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM="),
			},
		},
		{
			data: b64ToBytes("xL2OQM8Z7a5QloweIkbbBv45sxtX/j4/84h5HmqQxUE="),
			nonce: Nonce{
				Value:   56666,
				Payload: b64ToBytes("AACan9a7F7IwjPDPMmMy1sFaXag0hLflf/uquU9HrSo="),
			},
		},
	}

	for _, test := range tests {
		nonce, err := FindNonce(test.data)
		assert.Equal(t, test.nonce, *nonce)
		assert.NoError(t, err)
	}
}
