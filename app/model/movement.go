package model

import "time"

type Movement struct {
	Id          uint64    `json:"id"`
	CatId       uint64    `json:"cat_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      uint8     `json:"status"`
	UserId      uint64    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *Movement) TableName() string {
	return "iron_movements"
}
