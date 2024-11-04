package models

import "time"

type Product struct {
	ID         int64   `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	Price      float64 `json:"price" db:"price"`
	PercentOff float64 `json:"percentOff" db:"percent_off"`
	CreatedAt  int64   `json:"createdAt" db:"created_at"`
	UpdatedAt  int64   `json:"updatedAt" db:"updated_at"`
}

// BeforeSave will set the CreatedAt and UpdatedAt timestamps
func (p *Product) BeforeSave(firstTime bool) {
	currentTime := time.Now().Unix()
	if !firstTime {
		p.CreatedAt = currentTime
	}
	p.UpdatedAt = currentTime
}
