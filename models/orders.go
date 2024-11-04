package models

import "time"

type Orders struct {
	ID        int64 `json:"id" db:"id"`
	UserID    int64 `json:"userId" db:"user_id"`
	ProductID int64 `json:"productId" db:"product_id"`
	Quantity  int64 `json:"quantity" db:"quantity"`
	CreatedAt int64 `json:"createdAt" db:"created_at"`
	UpdatedAt int64 `json:"updatedAt" db:"updated_at"`
}

// BeforeSave will set the CreatedAt and UpdatedAt timestamps
func (o *Orders) BeforeSave(firstTime bool) {
	currentTime := time.Now().Unix()
	if !firstTime {
		o.CreatedAt = currentTime
	}
	o.UpdatedAt = currentTime
}
