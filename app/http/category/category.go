package category

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/lib/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/service"
	"net/http"
)

func Create(ctx *gin.Context) {
	var (
		req             request.CreateCategoryReq
		categoryService = service.NewCategoryService()
		id              uint64
		err             error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message1": err,
		})
		return
	}
	if id, err = categoryService.Create(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message2": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建类别成功!",
		"data":    id,
	})
}

func GetList(ctx *gin.Context) {
	var (
		categoryService = service.NewCategoryService()
		categories      []model.Category
		err             error
	)
	if categories, err = categoryService.GetList(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/api/categories [get]",
		"data":    categories,
	})
}
