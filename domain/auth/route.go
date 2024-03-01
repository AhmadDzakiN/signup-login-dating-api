package service

import (
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authService AuthService
}

func NewAuthRoute(authSvc AuthService) AuthRoute {
	return AuthRoute{authSvc}
}

func (a *AuthRoute) AuthRoute(r *gin.Engine) {
	r.POST("/register", a.authService.Register)
	r.POST("/login", a.authService.Login)
}
