// Package utils package utils
package utils

import (
	"crypto/rand"
	"math/big"
)

// RandomString RandomString
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		randomBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		s[i] = letters[randomBigInt.Int64()]
	}
	return string(s)
}

// RandomASCIILowerOctdigitsString -.
func RandomASCIILowerOctdigitsString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz01234567")

	s := make([]rune, n)
	for i := range s {
		randomBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		s[i] = letters[randomBigInt.Int64()]
	}
	return string(s)
}
