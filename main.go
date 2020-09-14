package main

import (
	"github.com/gin-gonic/gin"
	"iron/app/controllers/user"
	"net/http"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		// ping
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ping",
			})
		})
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
	r.Run()
}
