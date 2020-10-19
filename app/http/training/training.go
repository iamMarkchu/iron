package training

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/service"
	"net/http"
)

func Create(ctx *gin.Context) {
	var (
		req request.CreateTrainingReq
		trainingService = service.NewTrainingService()
		training *model.Training
		err error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message1": err,
		})
		return
	}
	if training, err = trainingService.Create(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message2": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "训练模块 创建训练",
		"data": training,
	})
}
