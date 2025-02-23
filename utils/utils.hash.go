package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

// GenerateSalt creates a random salt (s)
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// Generate Verifier (v) using salt (s) and password (p)
func GenerateVerifier(password, salt string) (string, error) {
	// Hash (password + salt) using SHA-256
	hashed := sha256.Sum256([]byte(password + salt))

	// Convert hash to a big integer
	v := new(big.Int).SetBytes(hashed[:])

	// Convert verifier to base64 for storage
	return base64.StdEncoding.EncodeToString(v.Bytes()), nil
}

// VerifyPassword compares input password against stored v and s
func VerifyPassword(password, salt, storedVerifier string) bool {
	// Recalculate verifier from input password and salt
	computedVerifier, err := GenerateVerifier(password, salt)
	if err != nil {
		return false
	}
	// Compare computed v with stored v
	return computedVerifier == storedVerifier
}
