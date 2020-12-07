package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
)

func GetOrders(c *gin.Context) {
	redis.NewRedis()
	list := redis.GetOrders()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": list,
	})
}
