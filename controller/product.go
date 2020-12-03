package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
)

func GetProduct(c *gin.Context) {
	id := c.Query("id")
	redis.NewRedis()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": redis.ProductInfo(id),
	})
}
