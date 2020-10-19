package request

import "time"

type CreateTrainingReq struct {
	TrainingDate    string                    `json:"training_date"`
	PlanId          int                       `json:"plan_id"`
	StartTime       time.Time                 `json:"start_time" time_format:"2006-01-02"`
	EndTime         time.Time                 `json:"end_time" time_format:"2006-01-02"`
	Description     string                    `json:"description"`
	TrainingDetails []*CreateTrainingLogItemReq `json:"plan_details"`
}

type CreateTrainingLogItemReq struct {
	PlanDetailId int   `json:"plan_detail_id"`
	Done         uint8 `json:"done"`
}

type GetTrainingsReq struct {
	UserId int `json:"user_id"`
}
