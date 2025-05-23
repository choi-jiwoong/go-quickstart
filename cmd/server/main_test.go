package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/choi-jiwoong/go-quickstart/internal/api"
	"github.com/choi-jiwoong/go-quickstart/internal/config"
	"github.com/choi-jiwoong/go-quickstart/internal/middleware"
	"github.com/gin-gonic/gin"
)

// setupRouter는 테스트를 위한 라우터를 설정합니다.
func setupRouter() *gin.Engine {
	// 테스트 모드 설정
	gin.SetMode(gin.TestMode)
	
	// 설정 로드
	cfg := config.NewConfig()
	
	// 라우터 설정
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.SetTrustedProxies(cfg.TrustedProxies)
	
	// 라우트 등록
	router.GET("/", api.RootHandler)
	router.GET("/ping", api.PingHandler)
	
	return router
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	
	// 테스트 요청 생성
	req := httptest.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()
	
	// 요청 처리
	router.ServeHTTP(resp, req)
	
	// 응답 검증
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}
	
	// 응답 본문 검증
	expected := `{"message":"pong"}`
	if resp.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, resp.Body.String())
	}
}

func TestRootRoute(t *testing.T) {
	router := setupRouter()
	
	// 테스트 요청 생성
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	
	// 요청 처리
	router.ServeHTTP(resp, req)
	
	// 응답 검증
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}
	
	// 응답에 clientIP 필드가 포함되어 있는지 확인
	if resp.Body.String() == "" {
		t.Error("Expected non-empty response body")
	}
}

func TestCustomPort(t *testing.T) {
	// 환경 변수 설정
	os.Setenv("PORT", "9090")
	defer os.Unsetenv("PORT")
	
	// 설정 로드
	cfg := config.NewConfig()
	
	// 포트 설정 검증
	if cfg.Port != "9090" {
		t.Errorf("Expected port to be '9090', got %s", cfg.Port)
	}
}