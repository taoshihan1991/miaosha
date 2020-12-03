package controller

import (
	"github.com/gin-gonic/gin"
)

//首页
func PageIndex(c *gin.Context) {
	if c.Request.RequestURI == "/favicon.ico" {
		return
	}
}
