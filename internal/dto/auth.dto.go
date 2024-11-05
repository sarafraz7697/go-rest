package dto

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"` // Username must be between 3 and 30 characters
	Email    string `json:"email" validate:"required,email"`           // Must be a valid email format
	Password string `json:"password" validate:"required,min=6"`        // Password must be at least 6 characters
}
