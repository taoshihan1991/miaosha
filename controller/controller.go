package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
	"log"
)

func checkUserLogin(c *gin.Context) (string, bool) {
	session := c.Query("session")
	info := redis.GetUserInfo(session)
	if session == "" || info == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error please login",
		})
		return "", false
	}

	return info, true
}
func limitIpFreq(c *gin.Context, timeWindow int64, count uint) bool {
	ip := c.ClientIP()
	key := "limit:" + ip
	//if !utils.LimitFreqSingle(key, count, timeWindow) {
	//	c.JSON(200, gin.H{
	//		"code": 400,
	//		"msg":  "error Current IP frequently visited",
	//	})
	//	return false
	//}
	if !redis.LimitFreqs(key, count, timeWindow) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Current IP frequently visited",
		})
		log.Println("key:" + key + " error Current IP frequently visited")
		return false
	}
	return true
}
