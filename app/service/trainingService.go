package service

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/core/store/orm"
)

type trainingService struct {

}

func (s *trainingService) Create(ctx *gin.Context, req request.CreateTrainingReq) (m *model.Training, err error) {
	var (
		o = orm.GetInstance()
	)
	m = &model.Training{}
	m.TrainingDate = req.TrainingDate
	m.PlanId = uint64(req.PlanId)
	m.StartTime = req.StartTime
	m.EndTime = req.EndTime
	m.Description = req.Description
	m.Status = model.StatusInit
	m.UserId = uint64(ctx.GetInt("userId"))
	o.Begin()
	o.Create(m)
	for _, item := range req.TrainingDetails {
		var tmp model.TrainingLog
		tmp.PlanDetailId = item.PlanDetailId
		tmp.Done = item.Done
		tmp.Status = model.StatusInit
		tmp.UserId = uint64(ctx.GetInt("userId"))
		o.Create(item)
		m.TrainingLogs = append(m.TrainingLogs, &tmp)
	}
	o.Commit()
	return
}

func NewTrainingService() *trainingService {
	return &trainingService{}
}
