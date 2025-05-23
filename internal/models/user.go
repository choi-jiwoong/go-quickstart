package models

import "time"

// User 모델은 사용자 정보를 나타냅니다.
type User struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Username  string     `json:"username" gorm:"size:50;unique;not null"`
	Email     string     `json:"email" gorm:"size:100;not null"`
	Password  string     `json:"-" gorm:"size:255;not null"` // JSON 응답에서 제외
	Role      string     `json:"role" gorm:"size:20;not null"`
}

// CreateUserRequest는 사용자 생성 요청을 나타냅니다.
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6,max=255"`
	Role     string `json:"role" binding:"required,oneof=USER ADMIN"`
}

// UpdateUserRequest는 사용자 업데이트 요청을 나타냅니다.
type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email,max=100"`
	Password string `json:"password" binding:"omitempty,min=6,max=255"`
	Role     string `json:"role" binding:"omitempty,oneof=USER ADMIN"`
}