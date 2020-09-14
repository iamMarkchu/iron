package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册接口
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/api/users/register [post]",
	})
}

// 用户登录接口
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "/api/users/login [post]",
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
