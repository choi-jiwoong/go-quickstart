package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choi-jiwoong/go-quickstart/internal/database"
	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 통합 테스트를 위한 설정
func setupIntegrationTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	
	// 인메모리 SQLite 데이터베이스 설정
	var err error
	database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("테스트 데이터베이스 연결 실패")
	}
	
	// 테이블 생성
	err = database.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("테스트 테이블 마이그레이션 실패")
	}
	
	// 테스트 데이터 생성
	users := []models.User{
		{
			Username: "testuser1",
			Email:    "test1@example.com",
			Password: "password1",
			Role:     "USER",
		},
		{
			Username: "testuser2",
			Email:    "test2@example.com",
			Password: "password2",
			Role:     "ADMIN",
		},
	}
	
	for _, user := range users {
		database.DB.Create(&user)
	}
	
	// 라우터 설정
	router := gin.New()
	router.GET("/users", GetUsers)
	router.GET("/user/:id", GetUser)
	router.POST("/user", CreateUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)
	
	return router
}

// TestUserCRUDIntegration은 사용자 CRUD 작업을 통합 테스트합니다.
func TestUserCRUDIntegration(t *testing.T) {
	router := setupIntegrationTest()
	
	// 1. 모든 사용자 조회
	t.Run("Get All Users", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users", nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var users []models.User
		err := json.Unmarshal(w.Body.Bytes(), &users)
		assert.NoError(t, err)
		assert.Len(t, users, 2)
	})
	
	// 2. 새 사용자 생성
	var newUserID int64
	t.Run("Create User", func(t *testing.T) {
		reqBody := models.CreateUserRequest{
			Username: "newuser",
			Email:    "new@example.com",
			Password: "newpassword",
			Role:     "USER",
		}
		body, _ := json.Marshal(reqBody)
		
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusCreated, w.Code)
		
		var user models.User
		err := json.Unmarshal(w.Body.Bytes(), &user)
		assert.NoError(t, err)
		assert.Equal(t, "newuser", user.Username)
		
		newUserID = user.ID
	})
	
	// 3. 특정 사용자 조회
	t.Run("Get User", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/user/%d", newUserID), nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var user models.User
		err := json.Unmarshal(w.Body.Bytes(), &user)
		assert.NoError(t, err)
		assert.Equal(t, "newuser", user.Username)
	})
	
	// 4. 사용자 업데이트
	t.Run("Update User", func(t *testing.T) {
		reqBody := models.UpdateUserRequest{
			Email: "updated@example.com",
			Role:  "ADMIN",
		}
		body, _ := json.Marshal(reqBody)
		
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/user/%d", newUserID), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var user models.User
		err := json.Unmarshal(w.Body.Bytes(), &user)
		assert.NoError(t, err)
		assert.Equal(t, "updated@example.com", user.Email)
		assert.Equal(t, "ADMIN", user.Role)
	})
	
	// 5. 사용자 삭제
	t.Run("Delete User", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/user/%d", newUserID), nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "사용자가 성공적으로 삭제되었습니다", response["message"])
		
		// 삭제 확인
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", fmt.Sprintf("/user/%d", newUserID), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}