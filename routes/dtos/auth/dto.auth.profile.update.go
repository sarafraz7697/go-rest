package auth

// UpdateProfileDTO defines the structure for updating a user's profile data.
type UpdateProfileDTO struct {
	Name        string `json:"name"`
	Family      string `json:"family"`
	Social_Name string `json:"social_name"`
	Email       string `json:"email"`
}
