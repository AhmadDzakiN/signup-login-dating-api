package user

import (
	"net/http"
	"signup-login-bumble/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func NewUserService(Db *gorm.DB) UserService {
	return UserService{Db: Db}
}

func (s *UserService) GetUserByID(ctx *gin.Context) {
	var user model.User

	userID := ctx.Param("user_id")
	getUserByID := s.Db.First(&user, "id = ?", userID)
	if getUserByID.Error != nil {
		if getUserByID.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": getUserByID.Error.Error()})
		return
	}

	response := model.GetUserResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Gender:   user.Gender,
		Age:      user.Age,
		Location: user.Location,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": response}})
}
