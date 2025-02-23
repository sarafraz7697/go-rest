package utils

import (
	"rest/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT creates a JWT token with user ID and role
func GenerateJWT(userId string, role string) (string, error) {
	// Get SECRET_KEY from environment variables
	secretKey := config.GetEnvOrFatal("SECRET_KEY")

	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseJWT validates and parses a token
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	// Get SECRET_KEY from environment variables
	secretKey := config.GetEnvOrFatal("SECRET_KEY")

	// Parse the token with the SECRET_KEY
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
