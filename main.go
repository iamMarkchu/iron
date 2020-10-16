package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/user"
	"github.com/iamMarkchu/iron/core"
	c "github.com/iamMarkchu/iron/core/config"
)

func main() {
	core.Init()
	// http
	r := gin.Default()
	// api router
	api := r.Group("/api")
	{
		// user模块
		userRouter := api.Group("/users")
		{
			userRouter.POST("/register", user.Register)
			userRouter.POST("/login", user.Login)
			userRouter.PUT("/:id", user.Edit)
			userRouter.POST("/forget", user.Forget)
			userRouter.POST("/reset", user.Reset)
		}
	}
	r.Run(c.GetInstance().GetString("common.webPort"))
}
