package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() (router *gin.Engine) {
	corsConfig := cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization", "Content-Length"},
	}

	corsOpt := cors.New(corsConfig)
	router = gin.Default()
	router.Use(corsOpt)

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return
}
