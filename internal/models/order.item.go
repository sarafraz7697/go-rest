package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // Store price at the time of order to prevent price changes
}
