package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
	"github.com/taoshihan1991/miaosha/utils"
	"log"
	"strconv"
	"time"
)

func GetProduct(c *gin.Context) {
	id := c.Query("id")
	redis.NewRedis()

	info := redis.ProductInfo(id)
	now := time.Now().UnixNano() / 1e6
	log.Println(info["saletime"], now)
	saleTime, err := strconv.Atoi(info["saletime"])
	if err != nil || int64(saleTime) < now {
		redis.SetProduct(id)
		info = redis.ProductInfo(id)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"product": info,
		},
	})
}
func GetKillUrl(c *gin.Context) {
	id := c.Query("id")
	redis.NewRedis()

	product := redis.ProductInfo(id)
	saletime, _ := strconv.Atoi(product["saletime"])
	nowTime := time.Now().UnixNano() / 1e6
	//时间
	if nowTime < int64(saletime) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Not started yet",
		})
		return
	}
	//库存
	storge, _ := strconv.Atoi(product["storge"])
	if storge <= 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Out of stock",
		})
		return
	}
	urlPath := fmt.Sprintf("product:%s,%d", id, time.Now().UnixNano())
	token := utils.Md5(urlPath)
	redis.SetStr(token, id, time.Second*10)
	url := "/seckill/" + token
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"url": url,
		},
	})
}
