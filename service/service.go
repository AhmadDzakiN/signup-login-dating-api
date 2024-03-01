package service

import (
	"net/http"
	"signup-login-bumble/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func NewService(Db *gorm.DB) Service {
	return Service{Db: Db}
}

func (s *Service) Register(ctx *gin.Context) {
	var userPayload model.UserRegisterPayload
	if err := ctx.ShouldBindJSON(&userPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if userPayload.Password != userPayload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password do not match"})
	}

	user := model.User{
		Username: userPayload.Username,
		Password: userPayload.Password,
		Fullname: userPayload.Fullname,
	}

	result := s.Db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (s *Service) Login(ctx *gin.Context) {
	var (
		loginPayload model.UserLoginPayload
		user         model.User
	)

	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := s.Db.First(&user, "username = ?", loginPayload.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}
