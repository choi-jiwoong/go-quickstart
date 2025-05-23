package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/choi-jiwoong/go-quickstart/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockUserRepository는 테스트를 위한 사용자 리포지토리 모의 객체입니다.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int64) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (models.User, error) {
	args := m.Called(username)
	return args.Get(0).(models.User), args.Error(1)
}

// 테스트 설정
func setupTest(t *testing.T) (*gin.Engine, *MockUserRepository) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockRepo := new(MockUserRepository)
	
	// 원래 함수 저장
	originalGetAllUsers := repository.GetAllUsers
	originalGetUserByID := repository.GetUserByID
	originalCreateUser := repository.CreateUser
	originalUpdateUser := repository.UpdateUser
	originalDeleteUser := repository.DeleteUser
	originalGetUserByUsername := repository.GetUserByUsername
	
	// 모의 함수로 대체
	repository.GetAllUsers = mockRepo.GetAllUsers
	repository.GetUserByID = mockRepo.GetUserByID
	repository.CreateUser = mockRepo.CreateUser
	repository.UpdateUser = mockRepo.UpdateUser
	repository.DeleteUser = mockRepo.DeleteUser
	repository.GetUserByUsername = mockRepo.GetUserByUsername
	
	t.Cleanup(func() {
		// 테스트 후 원래 함수 복원
		repository.GetAllUsers = originalGetAllUsers
		repository.GetUserByID = originalGetUserByID
		repository.CreateUser = originalCreateUser
		repository.UpdateUser = originalUpdateUser
		repository.DeleteUser = originalDeleteUser
		repository.GetUserByUsername = originalGetUserByUsername
	})
	
	return router, mockRepo
}

// TestGetUsers는 GetUsers 핸들러를 테스트합니다.
func TestGetUsers(t *testing.T) {
	router, mockRepo := setupTest(t)
	router.GET("/users", GetUsers)
	
	// 모의 데이터 설정
	users := []models.User{
		{ID: 1, Username: "user1", Email: "user1@example.com", Role: "USER"},
		{ID: 2, Username: "user2", Email: "user2@example.com", Role: "ADMIN"},
	}
	mockRepo.On("GetAllUsers").Return(users, nil)
	
	// 요청 실행
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)
	
	// 응답 검증
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "user1", response[0].Username)
	assert.Equal(t, "user2", response[1].Username)
	
	mockRepo.AssertExpectations(t)
}

// TestGetUser는 GetUser 핸들러를 테스트합니다.
func TestGetUser(t *testing.T) {
	router, mockRepo := setupTest(t)
	router.GET("/user/:id", GetUser)
	
	// 모의 데이터 설정
	user := models.User{ID: 1, Username: "user1", Email: "user1@example.com", Role: "USER"}
	mockRepo.On("GetUserByID", int64(1)).Return(user, nil)
	
	// 요청 실행
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/1", nil)
	router.ServeHTTP(w, req)
	
	// 응답 검증
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "user1", response.Username)
	
	mockRepo.AssertExpectations(t)
}

// TestCreateUser는 CreateUser 핸들러를 테스트합니다.
func TestCreateUser(t *testing.T) {
	router, mockRepo := setupTest(t)
	router.POST("/user", CreateUser)
	
	// 모의 데이터 설정
	mockRepo.On("GetUserByUsername", "newuser").Return(models.User{}, gorm.ErrRecordNotFound)
	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*models.User)
		user.ID = 1 // ID 설정
	})
	
	// 요청 데이터 생성
	reqBody := models.CreateUserRequest{
		Username: "newuser",
		Email:    "newuser@example.com",
		Password: "password123",
		Role:     "USER",
	}
	body, _ := json.Marshal(reqBody)
	
	// 요청 실행
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	
	// 응답 검증
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "newuser", response.Username)
	
	mockRepo.AssertExpectations(t)
}

// TestUpdateUser는 UpdateUser 핸들러를 테스트합니다.
func TestUpdateUser(t *testing.T) {
	router, mockRepo := setupTest(t)
	router.PUT("/user/:id", UpdateUser)
	
	// 모의 데이터 설정
	user := models.User{ID: 1, Username: "user1", Email: "user1@example.com", Role: "USER"}
	mockRepo.On("GetUserByID", int64(1)).Return(user, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)
	
	// 요청 데이터 생성
	reqBody := models.UpdateUserRequest{
		Email: "updated@example.com",
		Role:  "ADMIN",
	}
	body, _ := json.Marshal(reqBody)
	
	// 요청 실행
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	
	// 응답 검증
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "updated@example.com", response.Email)
	assert.Equal(t, "ADMIN", response.Role)
	
	mockRepo.AssertExpectations(t)
}

// TestDeleteUser는 DeleteUser 핸들러를 테스트합니다.
func TestDeleteUser(t *testing.T) {
	router, mockRepo := setupTest(t)
	router.DELETE("/user/:id", DeleteUser)
	
	// 모의 데이터 설정
	user := models.User{ID: 1, Username: "user1", Email: "user1@example.com", Role: "USER"}
	mockRepo.On("GetUserByID", int64(1)).Return(user, nil)
	mockRepo.On("DeleteUser", int64(1)).Return(nil)
	
	// 요청 실행
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	router.ServeHTTP(w, req)
	
	// 응답 검증
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "사용자가 성공적으로 삭제되었습니다", response["message"])
	
	mockRepo.AssertExpectations(t)
}