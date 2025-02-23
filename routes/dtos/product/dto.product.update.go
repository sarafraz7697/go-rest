package product

// UpdateProductDTO defines the structure for updating a product's details.
type UpdateProductDTO struct {
	Name        string  `json:"name,omitempty" validate:"omitempty"`
	Description string  `json:"description,omitempty" validate:"omitempty"`
	Price       float64 `json:"price,omitempty" validate:"omitempty,min=0"`
	Category    string  `json:"category,omitempty" validate:"omitempty"`
}
