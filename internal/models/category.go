package models

type Category struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	ParentID  *int64 `json:"parentId,omitempty" db:"parent_id"` // Optional
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
}
