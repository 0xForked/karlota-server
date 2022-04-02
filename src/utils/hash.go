package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

// Make returns the hash of the given string.
func (h Hash) Make(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)

	return string(bytes)
}

// Verify returns true if the given string matches the given hash.
func (h Hash) Verify(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))

	return err == nil
}
