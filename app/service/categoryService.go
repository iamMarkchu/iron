package service

import (
	"context"
	"errors"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/model/request"
	"github.com/iamMarkchu/iron/core/store/orm"
	"github.com/jinzhu/gorm"
)

type categoryService struct {
}

func NewCategoryService() *categoryService {
	return &categoryService{}
}

func (s *categoryService) Create(ctx context.Context, req request.CreateCategoryReq, uid uint64) (uint64, error) {
	var (
		o             = orm.GetInstance()
		categoryModel model.Category
		err           error
	)
	// 数据赋值 & 验证
	if categoryModel, err = s.valid(req, uid); err != nil {
		return 0, err
	}
	// 创建数据
	if err = o.Create(&categoryModel).Error; err != nil {
		return 0, err
	}
	return categoryModel.Id, err
}

func (s *categoryService) GetList(ctx context.Context) (categories []model.Category, err error) {
	if err = orm.GetInstance().Limit(100).Find(&categories).Error; err != nil {
		err = errors.New("查询错误")
	}
	return
}

func (s *categoryService) valid(req request.CreateCategoryReq, uid uint64) (cm model.Category, err error) {
	var (
		res model.Category
		o   = orm.GetInstance()
	)
	if len(req.Name) == 0 {
		err = errors.New("类别名称为空")
		return
	}
	if err = o.Table(res.TableName()).Where("name = ? and status = ?", req.Name, model.StatusInit).First(&res).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	}
	if res.Id > 0 {
		err = errors.New("存在同名类别名称")
		return
	}
	if req.ParentId > 0 {
		if err = o.Table(res.TableName()).Where("id = ? and status = ?", req.ParentId, model.StatusInit).First(&res).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return
			} else {
				err = errors.New("父类别不存在")
				return
			}
		}
	}
	// 赋值
	cm.Name = req.Name
	cm.ParentId = uint64(req.ParentId)
	cm.Description = model.ContentEmpty
	cm.Status = model.StatusInit
	cm.UserId = uid
	return
}

func (s *categoryService) GetTopCateList(ctx context.Context) ([]model.Category, error) {
	var (
		cm  model.Category
		cml []model.Category
		o   = orm.GetInstance()
		err error
	)
	if err = o.Table(cm.TableName()).Where("parent_id = ?", 0).Find(&cml).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return cml, nil
}
