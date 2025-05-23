package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger는 요청 로깅을 위한 미들웨어입니다.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 시작 시간
		startTime := time.Now()

		// 다음 핸들러 처리
		c.Next()

		// 요청 처리 시간 계산
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 요청 정보 로깅
		fmt.Printf("[%s] %s %s %d %s\n",
			endTime.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)
	}
}