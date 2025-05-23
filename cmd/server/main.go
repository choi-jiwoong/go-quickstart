package main

import (
	"fmt"
	"log"

	"github.com/choi-jiwoong/go-quickstart/internal/api"
	"github.com/choi-jiwoong/go-quickstart/internal/config"
	"github.com/choi-jiwoong/go-quickstart/internal/database"
	"github.com/choi-jiwoong/go-quickstart/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 설정 로드
	cfg := config.NewConfig()

	// 데이터베이스 초기화
	database.InitDB(cfg)

	// Gin 모드 설정
	gin.SetMode(cfg.GinMode)

	// 기본 라우터 설정
	router := gin.New()
	
	// 미들웨어 등록
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	
	// 신뢰할 수 있는 프록시 설정
	router.SetTrustedProxies(cfg.TrustedProxies)

	// 기존 라우트 등록
	router.GET("/", api.RootHandler)
	router.GET("/ping", api.PingHandler)
	
	// 인증 API 라우트 등록
	router.POST("/login", api.Login)
	
	// 인증이 필요한 API 그룹
	authGroup := router.Group("")
	authGroup.Use(middleware.RequireAuth())
	{
		// 사용자 조회 API
		authGroup.GET("/user/:id", api.GetUser)
		
		// 사용자 정보 업데이트 API
		authGroup.PUT("/user/:id", api.UpdateUser)
		
		// 관리자 전용 API 그룹
		adminGroup := authGroup.Group("")
		adminGroup.Use(middleware.RequireAdmin())
		{
			// 모든 사용자 목록 조회
			adminGroup.GET("/users", api.GetUsers)
			
			// 사용자 생성
			adminGroup.POST("/user", api.CreateUser)
			
			// 사용자 삭제
			adminGroup.DELETE("/user/:id", api.DeleteUser)
		}
	}

	// 서버 시작
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("서버가 %s 포트에서 시작됩니다...", cfg.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("서버 시작 실패: %v", err)
	}
}