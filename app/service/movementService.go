package service

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/lib/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/core/store/orm"
)

type movementService struct {
}

func (s *movementService) Create(ctx *gin.Context, req request.CreateMovementReq) (uint64, error) {
	var (
		o = orm.GetInstance()
		m model.Movement
	)
	m.Name = req.Name
	m.Description = req.Description
	m.CatId = uint64(req.CatId)
	m.UserId = uint64(ctx.GetInt("userId"))
	m.Status = model.StatusInit
	o.Create(&m)
	return m.Id, nil
}

func (s *movementService) GetList(ctx *gin.Context, req request.GetMovementsReq) (movements []model.Movement, err error) {
	var (
		o      = orm.GetInstance()
		offset int
	)
	if req.CatId > 0 {
		o = o.Where("cat_id = ?", req.CatId)
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}
	offset = (req.Page - 1) * req.Limit
	o.Offset(offset).Limit(req.Limit).Find(&movements)
	return
}

func NewMovementService() *movementService {
	return &movementService{}
}
