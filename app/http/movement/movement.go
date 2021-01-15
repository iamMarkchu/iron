package movement

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/model/request"
	"github.com/iamMarkchu/iron/app/service"
	"net/http"
)

func Create(ctx *gin.Context) {
	var (
		req             request.CreateMovementReq
		movementService = service.NewMovementService()
		id              uint64
		uid             uint64
		err             error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if uid = uint64(ctx.GetInt("userId")); uid == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "用户未登陆",
		})
		return
	}
	if id, err = movementService.Create(ctx, req, uid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建训练动作",
		"data":    id,
	})
}

func GetList(ctx *gin.Context) {
	var (
		req             request.GetMovementsReq
		movementService = service.NewMovementService()
		movements       []model.Movement
		err             error
	)
	if err = ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if movements, err = movementService.GetList(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取动作列表",
		"data":    movements,
	})
}
