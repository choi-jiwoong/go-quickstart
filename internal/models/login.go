package models

import "time"

// LoginRequest는 로그인 요청을 나타냅니다.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse는 로그인 응답을 나타냅니다.
type LoginResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token,omitempty"`
}

// LoginHistory는 로그인 시도 기록을 나타냅니다.
type LoginHistory struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	IPAddress string     `json:"ip_address" gorm:"size:50"`
	LoginTime *time.Time `json:"login_time" gorm:"not null"`
	Success   bool       `json:"success" gorm:"not null"`
	UserAgent string     `json:"user_agent" gorm:"size:255"`
	UserID    int64      `json:"user_id" gorm:"not null"`
	User      User       `json:"-" gorm:"foreignKey:UserID"`
}