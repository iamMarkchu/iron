package category

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/model/request"
	"github.com/iamMarkchu/iron/app/service"
	"net/http"
)

// @router /api/categories
// @description 创建类别
func Create(ctx *gin.Context) {
	var (
		req   request.CreateCategoryReq
		catId uint64
		uid   uint64
		err   error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
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
	if catId, err = service.CategorySer.Create(ctx, req, uid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建类别成功!",
		"data":    gin.H{
			"catId": catId,
		},
	})
}

func Edit(ctx *gin.Context) {
	var (
		req = request.CreateCategoryReq{}
		err error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {

	}
}

// @router /api/categories
func GetList(ctx *gin.Context) {
	var (
		categories      []model.Category
		err             error
	)
	if categories, err = service.CategorySer.GetList(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/api/categories [get]",
		"data":    categories,
	})
}

// @router /api/topCategories
// @description 获取顶级类别
func GetTopCate(ctx *gin.Context) {
	var (
		categories      []model.Category
		err             error
	)
	if categories, err = service.CategorySer.GetTopCateList(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/api/getTopCategories [get]",
		"data":    categories,
	})
}