package model

import "time"

type TrainingLog struct {
	PlanDetailId int `json:"plan_detail_id"`
	Done uint8 `json:"done"`
	Status int8 `json:"status"`
	UserId uint64 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *TrainingLog) TableName() string {
	return "iron_training_logs"
}