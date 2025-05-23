package api

import (
	"github.com/gin-gonic/gin"
)

// RootHandler는 루트 경로 요청을 처리합니다.
func RootHandler(c *gin.Context) {
	clientIP := c.ClientIP()
	c.JSON(200, gin.H{
		"clientIP": clientIP,
	})
}

// PingHandler는 /ping 경로 요청을 처리합니다.
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}