package auth

// RegisterDTO defines the structure for the registration data.
type RegisterDTO struct {
	Phone string `json:"phone" validate:"required"`
	// Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
