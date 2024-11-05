package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint       `gorm:"not null" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID"` // Relationship to User
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items"`
}
