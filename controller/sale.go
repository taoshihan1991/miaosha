package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
	"strconv"
	"strings"
	"time"
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
	username, bool := checkUserLogin(c)
	if !bool {
		return
	}
	id := redis.GetStr(p.Token)
	if id == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error token expire",
		})
		return
	}
	redis.DelKey(p.Token)
	//库存
	product := redis.ProductInfo(id)
	storge, _ := strconv.Atoi(product["storge"])
	if storge <= 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Out of stock",
		})
		return
	}
	if redis.OrderExist(username) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Order exist",
		})
		return
	}
	redis.PushRequestQueue(username + ":" + id)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success in queue",
	})
}
func GetProductQueueToOrder() {
	redis.NewRedis()
	for {
		item := redis.PopRequestQueue()
		if item != "" {
			locBool := redis.Lock("order_lock")

			itemSlice := strings.Split(item, ":")
			id := itemSlice[1]
			user := itemSlice[0]
			if id != "" && user != "" && locBool {
				if !redis.OrderExist(user) {
					storge := redis.DecProductStorge(id)
					if storge >= 0 {
						redis.InsertOrder(user, id)
					}
				}
			}
			redis.UnLock("order_lock")
		} else {
			time.Sleep(1 * time.Second)
		}

	}

}
