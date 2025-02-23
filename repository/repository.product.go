package repository

import (
	"context"
	"rest/config"
	"rest/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		collection: config.DB.Collection("products"),
	}
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FindAll products
func (r *ProductRepository) FindAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// FindByID product
func (r *ProductRepository) FindByID(ctx context.Context, id string) (*models.Product, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var product models.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update product
func (r *ProductRepository) Update(ctx context.Context, id string, updateData bson.M) (*models.Product, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": updateData}
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, id)
}

// Delete product
func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
