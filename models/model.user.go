package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Family     string             `bson:"family"`
	SocialName string             `bson:"social_name"`
	Email      string             `bson:"email"`

	Salt     string `json:"salt" bson:"salt"`         // salt
	Verifier string `json:"verifier" bson:"verifier"` // verifier

	Phone     string    `bson:"phone"`
	Role      string    `bson:"role"` // admin, user
	CreatedAt time.Time `bson:"created_at"`
}
