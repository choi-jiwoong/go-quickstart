package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

  router := gin.Default()
  router.SetTrustedProxies([]string{"192.168.1.2"})

  router.GET("/", func(c *gin.Context) {
    // 클라이언트가 192.168.1.2인 경우, X-Forwarded-For 헤더의 신뢰할 수 있는 부분을 사용하여 원래 클라이언트 IP를 추론합니다.
    // 그렇지 않으면, 단순히 직접 클라이언트 IP를 반환합니다.
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })


  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  router.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}