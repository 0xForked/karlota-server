package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

// Make returns the hash of the given string.
func (h Hash) Make(plain string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	
	return string(bytes)
}

// Verify returns true if the given string matches the given hash.
func (h Hash) Verify(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err == nil
}
