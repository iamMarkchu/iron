package plan

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/service"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context)  {
	var (
		req request.CreatePlanReq
		planService = service.NewPlanService()
		id uint64
		err error
	)
	if err = ctx.ShouldBindJSON(&req);err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	if id,err = planService.Create(ctx, req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建训练计划成功",
		"data": id,
	})
}

func GetList(ctx *gin.Context)  {
	var (
		req request.GetPlansReq
		planService = service.NewPlanService()
		plans []*model.Plan
		err error
	)
	if err = ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	if plans, err = planService.GetList(ctx, req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取训练计划列表",
		"data": plans,
	})
}

func Fetch(ctx *gin.Context)  {
	var (
		planId int
		planService = service.NewPlanService()
		plan model.Plan
		err error
	)
	sId := ctx.Param("id")
	if planId, err = strconv.Atoi(sId); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	if planId <= 0{
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	if plan, err = planService.Fetch(ctx, planId); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取单个训练计划列表",
		"data": plan,
	})
}