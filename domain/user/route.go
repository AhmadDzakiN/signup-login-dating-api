package user

import (
	"signup-login-bumble/middleware"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userService UserService
}

func NewUserRoute(userSvc UserService) UserRoute {
	return UserRoute{userSvc}
}

func (u *UserRoute) UserRoute(r *gin.Engine) {
	r.GET("/user/:user_id", middleware.BearerAuthCheck(), u.userService.GetUserByID)
}
