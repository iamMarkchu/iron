package request

type CreateTrainingReq struct {
	TrainingDate    string                      `json:"training_date"`
	PlanId          int                         `json:"plan_id"`
	StartTime       string                      `json:"start_time"`
	EndTime         string                      `json:"end_time"`
	Description     string                      `json:"description"`
	TrainingDetails []*CreateTrainingLogItemReq `json:"training_details"`
}

type CreateTrainingLogItemReq struct {
	PlanDetailId int   `json:"plan_detail_id"`
	Done         uint8 `json:"done"`
}

type GetTrainingsReq struct {
	UserId int `json:"user_id"`
}
