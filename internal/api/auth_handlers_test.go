package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// TestLogin은 로그인 핸들러를 테스트합니다.
func TestLogin(t *testing.T) {
	router, mockRepo := setupTest(t)
	router.POST("/login", Login)

	// 비밀번호 해시 생성
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	// 테스트 케이스
	tests := []struct {
		name           string
		requestBody    models.LoginRequest
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "로그인 성공",
			requestBody: models.LoginRequest{
				Username: "testuser",
				Password: "password123",
			},
			setupMock: func() {
				user := models.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: string(hashedPassword),
					Role:     "USER",
				}
				mockRepo.On("GetUserByUsername", "testuser").Return(user, nil)
				mockRepo.On("CreateLoginHistory", mock.AnythingOfType("*models.LoginHistory")).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"username":"testuser"`,
		},
		{
			name: "사용자가 존재하지 않음",
			requestBody: models.LoginRequest{
				Username: "nonexistent",
				Password: "password123",
			},
			setupMock: func() {
				mockRepo.On("GetUserByUsername", "nonexistent").Return(models.User{}, assert.AnError)
				mockRepo.On("CreateLoginHistory", mock.AnythingOfType("*models.LoginHistory")).Return(nil)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "사용자명 또는 비밀번호가 올바르지 않습니다",
		},
		{
			name: "비밀번호 불일치",
			requestBody: models.LoginRequest{
				Username: "testuser",
				Password: "wrongpassword",
			},
			setupMock: func() {
				user := models.User{
					ID:       1,
					Username: "testuser",
					Email:    "test@example.com",
					Password: string(hashedPassword),
					Role:     "USER",
				}
				mockRepo.On("GetUserByUsername", "testuser").Return(user, nil)
				mockRepo.On("CreateLoginHistory", mock.AnythingOfType("*models.LoginHistory")).Return(nil)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "사용자명 또는 비밀번호가 올바르지 않습니다",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 모의 객체 설정
			tt.setupMock()

			// 요청 데이터 생성
			body, _ := json.Marshal(tt.requestBody)

			// 요청 실행
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			// 응답 검증
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}