package model

import "time"

type Plan struct {
	Id          uint64        `json:"id"`
	PlanName    string        `json:"plan_name"`
	Status      int8          `json:"status"`
	UserId      uint64        `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	PlanDetails []*PlanDetail `json:"plan_details" gorm:"-"`
}

func (m *Plan) TableName() string {
	return "iron_plans"
}
