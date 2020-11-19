package user

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/model/request"
	"github.com/iamMarkchu/iron/app/service"
	"github.com/iamMarkchu/iron/core/auth/jwt"
	"net/http"
)

// 用户注册接口
// @params user_name
// @params password
// @params re_password
// @params mobile
// @params captcha
func Register(ctx *gin.Context) {
	var (
		req         request.RegisterReq
		userService = service.NewUserService()
		uid         int64
		err         error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if uid, err = userService.Register(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功!",
		"data":    uid,
	})
}

// 用户登录接口
func Login(ctx *gin.Context) {
	var (
		req         request.LoginReq
		userService = service.NewUserService()
		auth        jwt.Auth
		err         error
	)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if auth, err = userService.Login(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/api/users/login [post]",
		"data":    auth,
	})
}

// 修改用户信息
func Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/api/users/:id [put]",
		"id":      c.Param("id"),
	})
}

// 忘记密码
func Forget(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/api/users/forget [post]",
	})
}

// 重置密码
func Reset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/api/users/reset [post]",
	})
}

func WxLogin(ctx *gin.Context) {
	var (
		req         = new(request.WxLoginReq)
		userService = service.NewUserService()
		err         error
	)
	if err = ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userService.WxLogin(ctx, req.Code)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "/api/users/wxLogin [post]",
	})
}
