package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
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
