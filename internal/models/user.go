package models

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Salt     string `json:"salt"`     // To store the salt
	Verifier string `json:"verifier"` // To store the hashed password
}

// HashPassword hashes the user's password using bcrypt.
func (u *User) HashPassword(password string) error {
	salt := GenerateSalt() // Function to generate a random salt
	u.Salt = salt

	// Combine password and salt
	saltedPassword := append([]byte(password), []byte(salt)...)

	// Generate a bcrypt hash of the salted password
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Verifier = string(hashedPassword)
	return nil
}

// GenerateSalt generates a random salt.
func GenerateSalt() string {
	// Implement your salt generation logic, e.g., using a secure random generator
	const saltLength = 16
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		// handle error
	}
	return base64.StdEncoding.EncodeToString(salt)
}
