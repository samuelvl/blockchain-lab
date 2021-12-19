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

func TestInitTarget(t *testing.T) {
	target := initTarget(difficulty)
	fmt.Println(target)
}
