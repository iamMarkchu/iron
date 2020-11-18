package model

import "time"

type Training struct {
	Id           uint64         `json:"id"`
	TrainingDate string         `json:"training_date"`
	PlanId       uint64         `json:"plan_id"`
	StartTime    time.Time      `json:"start_time"`
	EndTime      time.Time      `json:"end_time"`
	Description  string         `json:"description"`
	Status       int8           `json:"status"`
	UserId       uint64         `json:"user_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	TrainingLogs []*TrainingLog `json:"training_logs" gorm:"-"`
}

func (m *Training) TableName() string {
	return "iron_trainings"
}
