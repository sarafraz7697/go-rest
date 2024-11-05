package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null" json:"user_id"`
	User        User        `gorm:"foreignKey:UserID" json:"user"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `gorm:"default:'pending'" json:"status"` // e.g., 'pending', 'completed', 'cancelled'
}
