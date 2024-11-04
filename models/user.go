package models

import "time"

type User struct {
	ID        int64   `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Family    string  `json:"family" db:"family"`
	Wallet    float64 `json:"wallet" db:"wallet"`
	CreatedAt int64   `json:"createdAt" db:"created_at"`
	UpdatedAt int64   `json:"updatedAt" db:"updated_at"`
}

// BeforeSave will set the CreatedAt and UpdatedAt timestamps
func (u *User) BeforeSave(firstTime bool) {
	currentTime := time.Now().Unix()
	if !firstTime {
		u.CreatedAt = currentTime
	}
	u.UpdatedAt = currentTime
}
