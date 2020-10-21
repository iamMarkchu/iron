package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/core/store/orm"
)

type planService struct {

}

func (s *planService) Create(ctx *gin.Context, req request.CreatePlanReq) (uint64, error) {
	var (
		o = orm.GetInstance()
		planModel model.Plan
		err error
	)
	tx := o.Begin()
	planModel.PlanName = req.PlanName
	planModel.Status = model.StatusInit
	planModel.UserId = uint64(ctx.GetInt("userId"))
	tx.Create(&planModel)
	for _, item := range req.PlanDetails {
		detailItemModel := model.PlanDetail{}
		detailItemModel.PlanId = planModel.Id
		detailItemModel.MovementId = uint64(item.MovementId)
		detailItemModel.Weight = item.Weight
		detailItemModel.Count = item.Count
		detailItemModel.Break = item.Break
		detailItemModel.UserId = planModel.UserId
		detailItemModel.Status = model.StatusInit
		if err = tx.Create(&detailItemModel).Error; err != nil {
			tx.Rollback()
			return 0, errors.New("保存训练详情失败")
		}
	}
	tx.Commit()
	return planModel.Id, nil
}

func (s *planService) GetList(ctx *gin.Context, req request.GetPlansReq) (plans []*model.Plan, err error) {
	var (
		o = orm.GetInstance()
	)
	o.Limit(10).Find(&plans)
	for _,item := range plans {
		o.Where("plan_id = ?", item.Id).Find(&item.PlanDetails)
	}
	return
}

func (s *planService) Fetch(ctx *gin.Context, id int) (plan model.Plan, err error){
	var (
		o = orm.GetInstance()
	)
	plan.Id = uint64(id)
	o.First(&plan)
	o.Where("plan_id =?", plan.Id).Find(&plan.PlanDetails)
	return
}

func NewPlanService() *planService {
	return &planService{}
}


