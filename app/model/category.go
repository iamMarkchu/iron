package model

import "time"

const (
	StatusInit = 1
)

type Category struct {
	Id uint64 `json:"id"`
	ParentId uint64 `json:"parent_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status int8 `json:"status"`
	UserId uint64 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Category) TableName() string {
	return "iron_categories"
}
