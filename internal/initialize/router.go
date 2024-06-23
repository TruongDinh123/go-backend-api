package initialize

import (
	"fmt"
	"net/http"
	"repodinh/go-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    fmt.Println("Before ---------------------------------------------------------------- AA")
    ctx.Next()
    fmt.Println("Alter ---------------------------------------------------------------- AA")
  }
}

func BB() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    fmt.Println("Before ---------------------------------------------------------------- BB")
    ctx.Next()
    fmt.Println("Alter ---------------------------------------------------------------- BB")
  }
}

func CC(ctx *gin.Context) {
  fmt.Println("Before ---------------------------------------------------------------- CC")
  ctx.Next()
  fmt.Println("Alter ---------------------------------------------------------------- CC")
}


func InitRouter() *gin.Engine {
  r := gin.Default()

  v1 := r.Group("/api/v1")
  {
    v1.GET("/ping", controller.NewUserController().GetUserByID)
  }

  return r
}

func Pong(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "pong",
  })
}