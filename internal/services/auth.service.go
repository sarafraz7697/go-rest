package services

import (
	"rest/internal/auth"
	"rest/internal/dto"
	"rest/internal/models"
	"rest/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(signInReq dto.SignInRequest) (string, error)
	Register(registerReq dto.RegisterRequest) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(registerReq dto.RegisterRequest) error {
	user := models.User{
		Username: registerReq.Username,
		Email:    registerReq.Email,
	}

	if err := user.HashPassword(registerReq.Password); err != nil {
		return err
	}

	return s.userRepo.Create(&user)
}

func (s *authService) SignIn(signInReq dto.SignInRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(signInReq.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Verifier), append([]byte(signInReq.Password), []byte(user.Salt)...)); err != nil {
		return "", err
	}

	token, err := auth.GenerateJWT(user.ID) // Assuming a JWT generation function
	return token, err
}
