package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	name := c.DefaultQuery("name", "dinhdev")

	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping!Controller" + name,
		"uid": uid,
		"user": []string{"dinh", "dev", "pro"},
	})
}