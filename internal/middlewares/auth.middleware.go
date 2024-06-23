package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			fmt.Println("invalid token: ", token)
			ctx.Abort()
			return
		}
	}
}