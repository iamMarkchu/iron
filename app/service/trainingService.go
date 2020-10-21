package service

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/core/store/orm"
	"time"
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
	m.StartTime, _ = time.Parse("2006-01-02 15:04:05", req.StartTime)
	m.EndTime,_ = time.Parse("2006-01-02 15:04:05", req.EndTime)
	m.Description = req.Description
	m.Status = model.StatusInit
	m.UserId = uint64(ctx.GetInt("userId"))
	tx := o.Begin()
	tx.Create(m)
	for _, item := range req.TrainingDetails {
		var tmp model.TrainingLog
		tmp.PlanDetailId = item.PlanDetailId
		tmp.Done = item.Done
		tmp.Status = model.StatusInit
		tmp.UserId = uint64(ctx.GetInt("userId"))
		err = tx.Create(&tmp).Error
		m.TrainingLogs = append(m.TrainingLogs, &tmp)
	}
	tx.Commit()
	return
}

func NewTrainingService() *trainingService {
	return &trainingService{}
}
