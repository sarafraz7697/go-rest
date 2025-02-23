package product

// CreateProductDTO defines the structure for creating a product.
type CreateProductDTO struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Stock       int     `json:"stock" validate:"required,min=0"`
	Category    string  `json:"category" validate:"required"`
}
