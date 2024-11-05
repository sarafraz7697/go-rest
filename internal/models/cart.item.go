package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null" json:"cart_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int     `json:"quantity"`
}
