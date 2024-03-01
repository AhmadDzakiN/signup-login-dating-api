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
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	if registerPayload.Password != registerPayload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Password do not match"})
		return
	}

	var checkExistUser model.User
	getExistUser := s.Db.First(&checkExistUser, "username = ?", registerPayload.Username)
	if getExistUser.Error != nil && getExistUser.Error != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": getExistUser.Error.Error()})
		return
	}

	if checkExistUser.ID != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": fmt.Sprintf("User with username %s already exist", registerPayload.Username)})
		return
	}

	hashedPassword, err := utils.HashPassword(registerPayload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err})
		return
	}

	user := model.User{
		Username: registerPayload.Username,
		Password: hashedPassword,
		Fullname: registerPayload.Fullname,
		Gender:   registerPayload.Gender,
		Location: registerPayload.Location,
		Age:      registerPayload.Age,
	}

	result := s.Db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": result.Error.Error()})
		return
	}

	response := model.GetUserResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Gender:   user.Gender,
		Age:      user.Age,
		Location: user.Location,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": response}})
}

func (s *AuthService) Login(ctx *gin.Context) {
	var (
		loginPayload model.UserLoginPayload
		user         model.User
	)

	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	getUser := s.Db.First(&user, "username = ?", loginPayload.Username)
	if getUser.Error != nil {
		if getUser.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "error": getUser.Error.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": getUser.Error.Error()})
		return
	}

	err := utils.CheckPassword(loginPayload.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Wrong password"})
		return
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Failed to create user token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"token": token}})
}
