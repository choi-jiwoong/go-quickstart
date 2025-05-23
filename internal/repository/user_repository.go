package repository

import (
	"github.com/choi-jiwoong/go-quickstart/internal/database"
	"github.com/choi-jiwoong/go-quickstart/internal/models"
)

// 함수 변수 선언 - 테스트에서 모킹하기 위함
var (
	GetAllUsers       = getAllUsers
	GetUserByID       = getUserByID
	CreateUser        = createUser
	UpdateUser        = updateUser
	DeleteUser        = deleteUser
	GetUserByUsername = getUserByUsername
)

// getAllUsers는 모든 사용자를 조회합니다.
func getAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

// getUserByID는 ID로 사용자를 조회합니다.
func getUserByID(id int64) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return user, result.Error
}

// createUser는 새 사용자를 생성합니다.
func createUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// updateUser는 사용자 정보를 업데이트합니다.
func updateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

// deleteUser는 사용자를 삭제합니다.
func deleteUser(id int64) error {
	return database.DB.Delete(&models.User{}, id).Error
}

// getUserByUsername은 사용자명으로 사용자를 조회합니다.
func getUserByUsername(username string) (models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}