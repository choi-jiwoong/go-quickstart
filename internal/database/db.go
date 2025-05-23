package database

import (
	"fmt"
	"log"

	"github.com/choi-jiwoong/go-quickstart/internal/config"
	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB는 데이터베이스 연결을 초기화합니다.
func InitDB(cfg *config.Config) {
	var err error
	
	// GORM 로거 설정
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	
	// 데이터베이스 연결
	dsn := cfg.GetDSN()
	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}
	
	fmt.Println("데이터베이스 연결 성공")
	
	// 모델 마이그레이션
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("마이그레이션 실패: %v", err)
	}
}