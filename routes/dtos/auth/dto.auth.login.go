package auth

// LoginDTO defines the structure for the login data.
type LoginDTO struct {
	Phone    string `json:"phone" validate:"required,phone"`
	Password string `json:"password" validate:"required,min=6"`
}
