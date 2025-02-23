package repository

import (
	"context"
	"rest/config"
	"rest/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	middleware "rest/middlewares"
	dtos "rest/routes/dtos/auth"
)

type AuthRepository struct {
	collection *mongo.Collection
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		collection: config.DB.Collection("users"),
	}
}

// Register User
func (repo *AuthRepository) Create(user *dtos.RegisterDTO, userID primitive.ObjectID, s, v string) (*mongo.InsertOneResult, error) {
	return repo.collection.InsertOne(context.TODO(), models.User{ID: userID, Phone: user.Phone, Verifier: v, Salt: s, Role: middleware.USER, CreatedAt: time.Now()})
}

// Find User By Phone
func (repo *AuthRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User

	// Query to find the user by phone number
	err := repo.collection.FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If no documents are found, return nil for the user
			return nil, nil
		}
		return nil, err // Return any other errors
	}
	return &user, nil
}

// Update Profile
func (repo *AuthRepository) UpdateProfile(id primitive.ObjectID, update *dtos.UpdateProfileDTO) error {
	_, err := repo.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}
