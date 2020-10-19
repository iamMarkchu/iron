package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/category"
	md "github.com/iamMarkchu/iron/app/http/middleware"
	"github.com/iamMarkchu/iron/app/http/movement"
	"github.com/iamMarkchu/iron/app/http/plan"
	"github.com/iamMarkchu/iron/app/http/user"
	"github.com/iamMarkchu/iron/core"
	c "github.com/iamMarkchu/iron/core/config"
)

func main() {
	core.Init()
	r := gin.Default()
	api := r.Group("/api")								// api接口前缀
	{
		api.POST("/register", user.Register)  					// 用户注册
		api.POST("/login", user.Login)        					// 用户登录
		api.Use(md.AuthRequired())
		{
			api.PUT("/users/:id", user.Edit)            // todo 用户信息修改
			api.POST("/users/forget", user.Forget)      // todo 忘记密码
			api.POST("/users/reset", user.Reset)        // todo 用户密码重置
			api.GET("/categories", category.GetList)    // done 训练类别模块 获取类别列表
			api.POST("/categories", category.Create)    // done 训练类别模块 创建类别
			api.POST("/movements", movement.Create)     // done 训练动作模块 创建训练动作
			api.GET("/movements", movement.GetList)     // done 训练动作模块 获取训练动作列表
			api.POST("/plans", plan.Create)             // done 训练计划模块 创建训练计划
			api.GET("/plans", plan.GetList)             // 训练计划模块 获取训练计划列表 eg: by user_id
			api.GET("/plans/:id", plan.Fetch)             // 训练计划模块 获取单个训练计划
		}
	}
	r.Run(c.GetInstance().GetString("common.webPort"))
}
