package main

import (
	"net/http"
	"signup-login-bumble/config"
	"signup-login-bumble/service"

	"github.com/gin-gonic/gin"
)

func setupRouter(svc service.Service) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/register", svc.Register)
	r.POST("/login", svc.Login)

	return r
}

func main() {
	config.LoadEnv()
	db, _ := config.ConnectDatabase()
	svc := service.NewService(db)

	r := setupRouter(svc)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
