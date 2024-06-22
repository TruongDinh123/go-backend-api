package routers

import (
	"net/http"
	"repodinh/go-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
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