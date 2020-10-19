package service

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/core/store/orm"
)

type categoryService struct {

}

func NewCategoryService() *categoryService {
	return &categoryService{}
}

func (s *categoryService) Create(ctx *gin.Context, req request.CreateCategoryReq) (uint64, error) {
	var (
		o = orm.GetInstance()
		categoryModel model.Category
	)
	categoryModel.Name = req.Name
	categoryModel.ParentId = uint64(req.ParentId)
	categoryModel.Description = " "
	categoryModel.Status = model.StatusInit
	categoryModel.UserId = uint64(ctx.GetInt("userId"))
	o.Create(&categoryModel)
	return categoryModel.Id, nil
}

func (s *categoryService) GetList(ctx *gin.Context) (categories []model.Category, err error) {
	var (
		o = orm.GetInstance()
	)
	o.Limit(10).Find(&categories)
	return
}
