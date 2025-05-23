package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	// 표준 출력을 캡처하기 위한 설정
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Gin 라우터 설정
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Logger())
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 테스트 요청 실행
	req := httptest.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// 표준 출력 복원 및 캡처된 출력 읽기
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 검증
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}

	// 로그 출력 검증
	if !strings.Contains(output, "GET") || !strings.Contains(output, "/test") || !strings.Contains(output, "200") {
		t.Errorf("Logger output doesn't contain expected information: %s", output)
	}
}