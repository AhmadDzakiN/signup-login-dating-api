package service

import (
	"fmt"
	"net/http"
	"signup-login-bumble/model"
	"signup-login-bumble/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService struct {
	Db *gorm.DB
}

func NewAuthService(Db *gorm.DB) AuthService {
	return AuthService{Db: Db}
}

func (s *AuthService) Register(ctx *gin.Context) {
	var registerPayload model.UserRegisterPayload
	if err := ctx.ShouldBindJSON(&registerPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if registerPayload.Password != registerPayload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password do not match"})
	}

	var checkExistUser model.User
	getExistUser := s.Db.First(&checkExistUser, "username = ?", registerPayload.Username)
	if getExistUser.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": getExistUser.Error.Error()})
		return
	}

	if checkExistUser.ID != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("User with username %s already exist", registerPayload.Username)})
		return
	}

	hashedPassword, err := utils.HashPassword(registerPayload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user := model.User{
		Username: registerPayload.Username,
		Password: hashedPassword,
		Fullname: registerPayload.Fullname,
	}

	result := s.Db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (s *AuthService) Login(ctx *gin.Context) {
	var (
		loginPayload model.UserLoginPayload
		user         model.User
	)

	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	getUser := s.Db.First(&user, "username = ?", loginPayload.Username)
	if getUser.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": getUser.Error.Error()})
		return
	}

	err := utils.CheckPassword(loginPayload.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}
