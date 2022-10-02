package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/go-gin-template/model"
)

func PingGet(c *gin.Context) {
	res := model.Ping{
		Message: "pong",
	}
	c.JSON(200, res)
}

func PingPost(c *gin.Context) {
	req := model.Ping{}
	c.BindJSON(&req)
	res := model.Ping{
		Message: req.Message,
	}
	c.JSON(200, res)
}
