package request

type CreateCategoryReq struct {
	ParentId int    `json:"parent_id"`
	Name     string `json:"name" binding:"required"`
}
