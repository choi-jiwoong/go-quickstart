package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/gin-gonic/gin"
)

// AuthUser는 인증된 사용자 정보를 저장하는 키입니다.
const AuthUser = "auth_user"

// RequireAuth는 인증이 필요한 엔드포인트에 대한 미들웨어입니다.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 헤더에서 토큰 추출
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "인증이 필요합니다"})
			c.Abort()
			return
		}

		// Bearer 토큰 형식 확인
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "잘못된 인증 형식입니다"})
			c.Abort()
			return
		}

		token := parts[1]
		
		// 실제 구현에서는 JWT 토큰 검증 필요
		// 여기서는 간단한 예시로 "admin-token"과 "user-token-{id}" 형식만 검증
		
		var user models.User
		
		if token == "admin-token" {
			// 관리자 토큰인 경우
			user = models.User{
				ID:       1,
				Username: "admin",
				Email:    "admin@example.com",
				Role:     "ADMIN",
			}
		} else if strings.HasPrefix(token, "user-token-") {
			// 사용자 토큰인 경우
			idStr := strings.TrimPrefix(token, "user-token-")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "잘못된 토큰입니다"})
				c.Abort()
				return
			}
			
			user = models.User{
				ID:       id,
				Username: "user" + idStr,
				Email:    "user" + idStr + "@example.com",
				Role:     "USER",
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 토큰입니다"})
			c.Abort()
			return
		}
		
		// 사용자 정보를 컨텍스트에 저장
		c.Set(AuthUser, user)
		c.Next()
	}
}

// RequireAdmin은 관리자 권한이 필요한 엔드포인트에 대한 미들웨어입니다.
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 인증된 사용자 정보 가져오기
		userInterface, exists := c.Get(AuthUser)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "인증이 필요합니다"})
			c.Abort()
			return
		}
		
		user, ok := userInterface.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "사용자 정보 처리 중 오류가 발생했습니다"})
			c.Abort()
			return
		}
		
		// 관리자 권한 확인
		if user.Role != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{"error": "관리자 권한이 필요합니다"})
			c.Abort()
			return
		}
		
		c.Next()
	}
}

// GetAuthUser는 컨텍스트에서 인증된 사용자 정보를 가져옵니다.
func GetAuthUser(c *gin.Context) (models.User, bool) {
	userInterface, exists := c.Get(AuthUser)
	if !exists {
		return models.User{}, false
	}
	
	user, ok := userInterface.(models.User)
	return user, ok
}