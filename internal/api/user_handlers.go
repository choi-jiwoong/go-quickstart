package api

import (
	"net/http"
	"strconv"

	"github.com/choi-jiwoong/go-quickstart/internal/middleware"
	"github.com/choi-jiwoong/go-quickstart/internal/models"
	"github.com/choi-jiwoong/go-quickstart/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers는 모든 사용자 목록을 반환합니다.
// 관리자만 접근 가능합니다.
func GetUsers(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "사용자 목록을 가져오는 중 오류가 발생했습니다",
		})
		return
	}
	
	c.JSON(http.StatusOK, users)
}

// GetUser는 특정 ID의 사용자 정보를 반환합니다.
// 관리자는 모든 사용자 정보를 볼 수 있고, 일반 사용자는 자신의 정보만 볼 수 있습니다.
func GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 사용자 ID 형식입니다",
		})
		return
	}
	
	// 인증된 사용자 정보 가져오기
	authUser, _ := middleware.GetAuthUser(c)
	
	// 권한 확인: 관리자가 아니고 자신의 정보가 아닌 경우 접근 거부
	if authUser.Role != "ADMIN" && authUser.ID != id {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "다른 사용자의 정보에 접근할 권한이 없습니다",
		})
		return
	}
	
	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "사용자를 찾을 수 없습니다",
		})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// CreateUser는 새 사용자를 생성합니다.
// 관리자만 접근 가능합니다.
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 요청 형식입니다: " + err.Error(),
		})
		return
	}
	
	// 사용자명 중복 확인
	_, err := repository.GetUserByUsername(req.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "이미 사용 중인 사용자명입니다",
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "사용자 확인 중 오류가 발생했습니다",
		})
		return
	}
	
	// 새 사용자 생성
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // 실제 구현에서는 비밀번호 해싱 필요
		Role:     req.Role,
	}
	
	if err := repository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "사용자 생성 중 오류가 발생했습니다",
		})
		return
	}
	
	c.JSON(http.StatusCreated, user)
}

// UpdateUser는 사용자 정보를 업데이트합니다.
// 관리자는 모든 사용자 정보를 수정할 수 있고, 일반 사용자는 자신의 정보만 수정할 수 있습니다.
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 사용자 ID 형식입니다",
		})
		return
	}
	
	// 인증된 사용자 정보 가져오기
	authUser, _ := middleware.GetAuthUser(c)
	
	// 권한 확인: 관리자가 아니고 자신의 정보가 아닌 경우 접근 거부
	if authUser.Role != "ADMIN" && authUser.ID != id {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "다른 사용자의 정보를 수정할 권한이 없습니다",
		})
		return
	}
	
	// 기존 사용자 조회
	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "사용자를 찾을 수 없습니다",
		})
		return
	}
	
	// 요청 바인딩
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 요청 형식입니다: " + err.Error(),
		})
		return
	}
	
	// 일반 사용자는 역할을 변경할 수 없음
	if authUser.Role != "ADMIN" && req.Role != "" && req.Role != user.Role {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "역할을 변경할 권한이 없습니다",
		})
		return
	}
	
	// 사용자명 변경 시 중복 확인
	if req.Username != "" && req.Username != user.Username {
		_, err := repository.GetUserByUsername(req.Username)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "이미 사용 중인 사용자명입니다",
			})
			return
		} else if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "사용자 확인 중 오류가 발생했습니다",
			})
			return
		}
		user.Username = req.Username
	}
	
	// 필드 업데이트
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password // 실제 구현에서는 비밀번호 해싱 필요
	}
	if req.Role != "" && authUser.Role == "ADMIN" {
		user.Role = req.Role
	}
	
	// 사용자 정보 저장
	if err := repository.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "사용자 업데이트 중 오류가 발생했습니다",
		})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// DeleteUser는 사용자를 삭제합니다.
// 관리자만 접근 가능합니다.
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "잘못된 사용자 ID 형식입니다",
		})
		return
	}
	
	// 사용자 존재 여부 확인
	_, err = repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "사용자를 찾을 수 없습니다",
		})
		return
	}
	
	// 사용자 삭제
	if err := repository.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "사용자 삭제 중 오류가 발생했습니다",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "사용자가 성공적으로 삭제되었습니다",
	})
}