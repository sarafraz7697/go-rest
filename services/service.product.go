package services

import (
	"context"
	"rest/models"
	"rest/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct service
func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.Create(ctx, product)
}

// GetAllProducts service
func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.FindAll(ctx)
}

// GetProductByID service
func (s *ProductService) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.FindByID(ctx, id)
}

// UpdateProduct service
func (s *ProductService) UpdateProduct(ctx context.Context, id string, updateData bson.M) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.Update(ctx, id, updateData)
}

// DeleteProduct service
func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.Delete(ctx, id)
}
