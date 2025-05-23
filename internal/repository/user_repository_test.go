package repository

import (
	"testing"
	"time"

	"github.com/choi-jiwoong/go-quickstart/internal/database"
	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 테스트용 데이터베이스 설정
func setupTestDB() {
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
}

// 테스트 데이터 생성
func createTestUsers() []models.User {
	now := time.Now()
	users := []models.User{
		{
			Username:  "testuser1",
			Email:     "test1@example.com",
			Password:  "password1",
			Role:      "USER",
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			Username:  "testuser2",
			Email:     "test2@example.com",
			Password:  "password2",
			Role:      "ADMIN",
			CreatedAt: &now,
			UpdatedAt: &now,
		},
	}
	
	for i := range users {
		database.DB.Create(&users[i])
	}
	
	return users
}

// 테스트 데이터 정리
func cleanupTestData() {
	database.DB.Exec("DELETE FROM users")
}

// TestGetAllUsers는 GetAllUsers 함수를 테스트합니다.
func TestGetAllUsers(t *testing.T) {
	setupTestDB()
	testUsers := createTestUsers()
	defer cleanupTestData()
	
	users, err := GetAllUsers()
	
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, testUsers[0].Username, users[0].Username)
	assert.Equal(t, testUsers[1].Username, users[1].Username)
}

// TestGetUserByID는 GetUserByID 함수를 테스트합니다.
func TestGetUserByID(t *testing.T) {
	setupTestDB()
	testUsers := createTestUsers()
	defer cleanupTestData()
	
	user, err := GetUserByID(testUsers[0].ID)
	
	assert.NoError(t, err)
	assert.Equal(t, testUsers[0].Username, user.Username)
	assert.Equal(t, testUsers[0].Email, user.Email)
}

// TestCreateUser는 CreateUser 함수를 테스트합니다.
func TestCreateUser(t *testing.T) {
	setupTestDB()
	defer cleanupTestData()
	
	newUser := models.User{
		Username: "newuser",
		Email:    "new@example.com",
		Password: "newpassword",
		Role:     "USER",
	}
	
	err := CreateUser(&newUser)
	assert.NoError(t, err)
	assert.NotZero(t, newUser.ID)
	
	// 데이터베이스에서 사용자 조회
	var savedUser models.User
	database.DB.First(&savedUser, newUser.ID)
	
	assert.Equal(t, newUser.Username, savedUser.Username)
	assert.Equal(t, newUser.Email, savedUser.Email)
}

// TestUpdateUser는 UpdateUser 함수를 테스트합니다.
func TestUpdateUser(t *testing.T) {
	setupTestDB()
	testUsers := createTestUsers()
	defer cleanupTestData()
	
	user := testUsers[0]
	user.Email = "updated@example.com"
	user.Role = "ADMIN"
	
	err := UpdateUser(&user)
	assert.NoError(t, err)
	
	// 데이터베이스에서 사용자 조회
	var updatedUser models.User
	database.DB.First(&updatedUser, user.ID)
	
	assert.Equal(t, "updated@example.com", updatedUser.Email)
	assert.Equal(t, "ADMIN", updatedUser.Role)
}

// TestDeleteUser는 DeleteUser 함수를 테스트합니다.
func TestDeleteUser(t *testing.T) {
	setupTestDB()
	testUsers := createTestUsers()
	defer cleanupTestData()
	
	err := DeleteUser(testUsers[0].ID)
	assert.NoError(t, err)
	
	// 데이터베이스에서 사용자 조회
	var count int64
	database.DB.Model(&models.User{}).Where("id = ?", testUsers[0].ID).Count(&count)
	
	assert.Equal(t, int64(0), count)
}

// TestGetUserByUsername은 GetUserByUsername 함수를 테스트합니다.
func TestGetUserByUsername(t *testing.T) {
	setupTestDB()
	testUsers := createTestUsers()
	defer cleanupTestData()
	
	user, err := GetUserByUsername(testUsers[0].Username)
	
	assert.NoError(t, err)
	assert.Equal(t, testUsers[0].Username, user.Username)
	assert.Equal(t, testUsers[0].Email, user.Email)
}