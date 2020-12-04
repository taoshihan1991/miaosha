package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetTimestamp(c *gin.Context) {
	nowTime := time.Now()
	timestamp := nowTime.UnixNano() / 1e6
	js := fmt.Sprintf("var nowTimestamp=%d;", timestamp)
	c.Writer.Write([]byte(js))
}
