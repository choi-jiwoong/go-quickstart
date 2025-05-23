package repository

import (
	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/stretchr/testify/mock"
)

// 원래 함수 타입 정의
var (
	GetAllUsersFunc      func() ([]models.User, error)
	GetUserByIDFunc      func(id int64) (models.User, error)
	CreateUserFunc       func(user *models.User) error
	UpdateUserFunc       func(user *models.User) error
	DeleteUserFunc       func(id int64) error
	GetUserByUsernameFunc func(username string) (models.User, error)
)

// 테스트를 위한 모의 함수 설정
func SetupMockRepository() {
	GetAllUsersFunc = func() ([]models.User, error) {
		return mockRepo.GetAllUsers()
	}
	GetUserByIDFunc = func(id int64) (models.User, error) {
		return mockRepo.GetUserByID(id)
	}
	CreateUserFunc = func(user *models.User) error {
		return mockRepo.CreateUser(user)
	}
	UpdateUserFunc = func(user *models.User) error {
		return mockRepo.UpdateUser(user)
	}
	DeleteUserFunc = func(id int64) error {
		return mockRepo.DeleteUser(id)
	}
	GetUserByUsernameFunc = func(username string) (models.User, error) {
		return mockRepo.GetUserByUsername(username)
	}
}

// 원래 함수 복원
func RestoreRepository() {
	GetAllUsersFunc = nil
	GetUserByIDFunc = nil
	CreateUserFunc = nil
	UpdateUserFunc = nil
	DeleteUserFunc = nil
	GetUserByUsernameFunc = nil
}

// MockRepository는 테스트를 위한 리포지토리 모의 객체입니다.
type MockRepository struct {
	mock.Mock
}

var mockRepo *MockRepository

// InitMockRepository는 모의 리포지토리를 초기화합니다.
func InitMockRepository() *MockRepository {
	mockRepo = new(MockRepository)
	SetupMockRepository()
	return mockRepo
}

func (m *MockRepository) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockRepository) GetUserByID(id int64) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) DeleteUser(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) GetUserByUsername(username string) (models.User, error) {
	args := m.Called(username)
	return args.Get(0).(models.User), args.Error(1)
}