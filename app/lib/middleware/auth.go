package middleware

import (
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/core/auth/jwt"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			token   string
			userId  int
			isValid bool
			err     error
		)
		if token, err = request.AuthorizationHeaderExtractor.ExtractToken(ctx.Request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "缺少token",
			})
			return
		}
		if userId, isValid = jwt.CheckToken(token); !isValid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "token验证失败",
			})
			return
		}
		ctx.Set("userId", userId)
	}
}
