package services

import (
	"errors"
	. "rest/middlewares"
	"rest/repository"
	"rest/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"

	dtos "rest/routes/dtos/auth"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Register(dto *dtos.RegisterDTO) (string, error) {
	// Check if the phone is already registered
	existingUser, err := s.repo.FindByPhone(dto.Phone)
	if err != nil {
		return "", err
	}

	if existingUser != nil && !existingUser.ID.IsZero() {
		return "", errors.New("phone already registered")
	}

	// Create a new ObjectID for the user
	userID := primitive.NewObjectID()

	// Generate a unique salt
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		return "", err
	}

	// Hash password with salt to create a verifier
	verifier, err := utils.GenerateVerifier(dto.Password, salt)
	if err != nil {
		return "", err
	}

	// Insert the new user into the database
	if _, err := s.repo.Create(dto, userID, salt, verifier); err != nil {
		return "", err
	}

	// Generate JWT token after registration
	token, err := utils.GenerateJWT(userID.Hex(), ADMIN)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login User (verify credentials and return JWT)
func (s *AuthService) Login(phone, password string) (string, error) {
	user, err := s.repo.FindByPhone(phone)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Verify password using stored salt (s) and verifier (v)
	if !utils.VerifyPassword(password, user.Salt, user.Verifier) {
		return "", errors.New("invalid phone or password")
	}

	token, _ := utils.GenerateJWT(user.ID.Hex(), user.Role)
	return token, nil
}

func (s *AuthService) UpdateProfile(id primitive.ObjectID, updateData *dtos.UpdateProfileDTO) error {
	return s.repo.UpdateProfile(id, updateData)
}
