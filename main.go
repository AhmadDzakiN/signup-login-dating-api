package main

import (
	"signup-login-bumble/config"
	auth "signup-login-bumble/domain/auth"
	"signup-login-bumble/domain/user"
	"signup-login-bumble/router"
)

func main() {
	config.LoadEnv()
	db, err := config.ConnectDatabase()
	if err != nil {
		return
	}

	router := router.Router()

	authSvc := auth.NewAuthService(db)
	userSvc := user.NewUserService(db)

	authRoute := auth.NewAuthRoute(authSvc)
	userRoute := user.NewUserRoute(userSvc)

	authRoute.AuthRoute(router)
	userRoute.UserRoute(router)

	router.Run(":8080")
}
