package repository

import (
	"time"

	"github.com/choi-jiwoong/go-quickstart/internal/database"
	"github.com/choi-jiwoong/go-quickstart/internal/models"
)

// 함수 변수 선언 - 테스트에서 모킹하기 위함
var (
	CreateLoginHistory = createLoginHistory
	GetLoginHistories  = getLoginHistories
)

// createLoginHistory는 로그인 시도 기록을 저장합니다.
func createLoginHistory(history *models.LoginHistory) error {
	if history.LoginTime == nil {
		now := time.Now()
		history.LoginTime = &now
	}
	return database.DB.Create(history).Error
}

// getLoginHistories는 특정 사용자의 로그인 기록을 조회합니다.
func getLoginHistories(userID int64) ([]models.LoginHistory, error) {
	var histories []models.LoginHistory
	result := database.DB.Where("user_id = ?", userID).Find(&histories)
	return histories, result.Error
}