// Package hash provides utility functions for hashing and comparing strings using bcrypt.
// If you want to use pure Go implementations of hashing algorithms, consider using the "crypto" package directly.
package hash

import (
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

// String hashes the given string using bcrypt with a cost of 10.
func String(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// MustString is a helper function that wraps function String and panics if an error occurs.
func MustString(str string) string {
	return lo.Must(String(str))
}

// Compare compares a plaintext string with a bcrypt hashed string and returns true if they match.
func Compare(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
