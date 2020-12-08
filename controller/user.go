package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/miaosha/redis"
	"github.com/taoshihan1991/miaosha/utils"
)

func GetUserInfo(c *gin.Context) {
	session := c.Query("session")
	redis.NewRedis()
	info := redis.GetUserInfo(session)
	if info == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error please login",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"username": info,
		},
	})
}
func PostUserInfo(c *gin.Context) {
	name := c.PostForm("name")
	session := utils.Md5(name)
	redis.NewRedis()
	info := redis.GetUserInfo(session)
	if info != "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error username exist",
		})
		return
	}
	redis.SetUserInfo(session, name)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": session,
	})
}
