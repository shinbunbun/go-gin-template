package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/go-gin-template/model"
)

func Ping(c *gin.Context) {
	res := model.Ping{
		Message: "pong",
	}
	c.JSON(200, res)
}
