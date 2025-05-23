package main

import (
	"fmt"
	"log"

	"github.com/choi-jiwoong/go-quickstart/internal/api"
	"github.com/choi-jiwoong/go-quickstart/internal/config"
	"github.com/choi-jiwoong/go-quickstart/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 설정 로드
	cfg := config.NewConfig()

	// Gin 모드 설정
	gin.SetMode(cfg.GinMode)

	// 기본 라우터 설정
	router := gin.New()
	
	// 미들웨어 등록
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	
	// 신뢰할 수 있는 프록시 설정
	router.SetTrustedProxies(cfg.TrustedProxies)

	// 라우트 등록
	router.GET("/", api.RootHandler)
	router.GET("/ping", api.PingHandler)

	// 서버 시작
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("서버가 %s 포트에서 시작됩니다...", cfg.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("서버 시작 실패: %v", err)
	}
}