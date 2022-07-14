package services

import (
	"github.com/multiformats/go-multihash"
)

// ETag computes a strong HTTP ETag string for a block of bytes.
func ETag(b []byte) (string, error) {
	mh, err := multihash.Sum(b, multihash.SHA1, -1)
	if err != nil {
		return "", err
	}
	return mh.B58String(), nil
}
