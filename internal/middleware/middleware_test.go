package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choi-jiwoong/go-quickstart/internal/middleware"
	"github.com/gin-gonic/gin"
)

// 미들웨어 체인 테스트
func TestMiddlewareChain(t *testing.T) {
	// Gin 테스트 모드 설정
	gin.SetMode(gin.TestMode)
	
	// 테스트용 라우터 설정
	router := gin.New()
	
	// 미들웨어 등록
	router.Use(middleware.Logger())
	
	// 테스트용 핸들러 등록
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	
	// 테스트 요청 실행
	req := httptest.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	
	// 응답 검증
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}
}

// 미들웨어 성능 테스트
func BenchmarkLogger(b *testing.B) {
	// Gin 테스트 모드 설정
	gin.SetMode(gin.TestMode)
	
	// 테스트용 라우터 설정
	router := gin.New()
	router.Use(middleware.Logger())
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	
	// 벤치마크 실행
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
	}
}