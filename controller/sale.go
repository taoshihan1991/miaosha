package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
)

type Product struct {
	Token string `uri:"token" binding:"required"`
}

func PostSale(c *gin.Context) {
	var p Product
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error no token",
		})
		return
	}
	redis.NewRedis()
	id := redis.GetStr(p.Token)
	if id == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error token expire",
		})
		return
	}
	redis.DelKey(p.Token)

	redis.Lock("order_lock")
	storge := redis.DecProductStorge(id)
	if storge >= 0 {
		redis.InsertOrder("taoshihan", id)
	}
	redis.UnLock("order_lock")

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
