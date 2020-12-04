package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
	"io"
	"strconv"
	"time"
)

func GetProduct(c *gin.Context) {
	id := c.Query("id")
	redis.NewRedis()

	h := md5.New()
	io.WriteString(h, fmt.Sprintf("product:%s,%d", id, time.Now().UnixNano()))
	token := string(h.Sum(nil))
	//redis.SetStr(token, 1, time.Hour*24)
	redis.SetProduct(id)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"token":   token,
			"product": redis.ProductInfo(id),
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
}
