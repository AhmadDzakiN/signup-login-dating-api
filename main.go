package main

import (
	"log"
	"signup-login-bumble/config"
	auth "signup-login-bumble/domain/auth"
	"signup-login-bumble/router"
)

func main() {
	config.LoadEnv()
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
		return
	}

	router := router.Router()

	authSvc := auth.NewAuthService(db)
	authRoute := auth.NewAuthRoute(authSvc)
	authRoute.AuthRoute(router)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
