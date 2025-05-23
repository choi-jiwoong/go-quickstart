package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/choi-jiwoong/go-quickstart/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login은 사용자 로그인을 처리합니다.
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 요청 형식입니다: " + err.Error(),
		})
		return
	}

	// 사용자 조회
	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		// 로그인 실패 기록
		recordLoginAttempt(c, 0, false)
		
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "사용자명 또는 비밀번호가 올바르지 않습니다",
		})
		return
	}

	// 비밀번호 검증
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// 로그인 실패 기록
		recordLoginAttempt(c, user.ID, false)
		
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "사용자명 또는 비밀번호가 올바르지 않습니다",
		})
		return
	}

	// 로그인 성공 기록
	recordLoginAttempt(c, user.ID, true)

	// 토큰 생성 (실제 구현에서는 JWT 토큰 생성 필요)
	var token string
	if user.Role == "ADMIN" {
		token = "admin-token"
	} else {
		token = fmt.Sprintf("user-token-%d", user.ID)
	}

	// 응답 생성
	response := models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	c.JSON(http.StatusOK, response)
}

// recordLoginAttempt는 로그인 시도를 기록합니다.
func recordLoginAttempt(c *gin.Context, userID int64, success bool) {
	now := time.Now()
	history := models.LoginHistory{
		IPAddress: c.ClientIP(),
		LoginTime: &now,
		Success:   success,
		UserAgent: c.Request.UserAgent(),
		UserID:    userID,
	}

	// 비동기적으로 로그인 기록 저장
	go func() {
		repository.CreateLoginHistory(&history)
	}()
}