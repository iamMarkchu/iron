package request

type CreatePlanReq struct {
	PlanName    string                    `json:"plan_name" binding:"required"`
	PlanDetails []CreatePlanDetailItemReq `json:"plan_details"`
}

type CreatePlanDetailItemReq struct {
	MovementId int `json:"movement_id" binding:"required"`
	Weight     int `json:"weight" binding:"required"`
	Count      int `json:"count" binding:"required"`
	Break      int `json:"break" binding:"required"`
}

type GetPlansReq struct {
	UserId int `json:"user_id"`
}
