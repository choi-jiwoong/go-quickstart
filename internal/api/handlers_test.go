package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRootHandler(t *testing.T) {
	// Gin 테스트 모드 설정
	gin.SetMode(gin.TestMode)
	
	// 테스트용 컨텍스트 생성
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	// 테스트용 요청 설정
	req := httptest.NewRequest("GET", "/", nil)
	c.Request = req
	
	// 핸들러 실행
	RootHandler(c)
	
	// 응답 검증
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	
	// 응답 본문 파싱
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	
	// clientIP 필드 존재 확인
	if _, exists := response["clientIP"]; !exists {
		t.Error("Expected 'clientIP' field in response")
	}
}

func TestPingHandler(t *testing.T) {
	// Gin 테스트 모드 설정
	gin.SetMode(gin.TestMode)
	
	// 테스트용 컨텍스트 생성
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	// 테스트용 요청 설정
	req := httptest.NewRequest("GET", "/ping", nil)
	c.Request = req
	
	// 핸들러 실행
	PingHandler(c)
	
	// 응답 검증
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	
	// 응답 본문 파싱
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	
	// message 필드 검증
	if message, exists := response["message"]; !exists || message != "pong" {
		t.Errorf("Expected message to be 'pong', got %v", message)
	}
}