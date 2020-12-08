package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
	"strconv"
)

func GetOrders(c *gin.Context) {
	start := c.Query("start")
	stop := c.Query("stop")
	startNum, _ := strconv.ParseInt(start, 10, 64)
	stopNum, _ := strconv.ParseInt(stop, 10, 64)
	if stopNum == 0 {
		stopNum = 10
	}
	redis.NewRedis()
	list := redis.GetOrders(startNum, stopNum)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": list,
	})
}
