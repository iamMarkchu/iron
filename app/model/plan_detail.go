package model

import "time"

type PlanDetail struct {
	Id uint64  `json:"id"`
	PlanId uint64 `json:"plan_id"`
	MovementId uint64 `json:"movement_id"`
	Weight int `json:"weight"`
	Count int `json:"count"`
	Break int `json:"break"`
	Status int8 `json:"status"`
	UserId uint64 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *PlanDetail) TableName() string {
	return "iron_plan_details"
}
