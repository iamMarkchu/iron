package request

type CreateMovementReq struct {
	CatId       int    `json:"cat_id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type GetMovementsReq struct {
	CatId int `form:"cat_id"`
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
