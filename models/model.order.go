package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	UserID     primitive.ObjectID   `bson:"user_id"`
	Products   []primitive.ObjectID `bson:"products"`
	TotalPrice float64              `bson:"total_price"`
	Status     string               `bson:"status"` // pending, shipped, delivered
	CreatedAt  time.Time            `bson:"created_at"`
}
