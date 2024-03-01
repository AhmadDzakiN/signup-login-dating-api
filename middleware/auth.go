package middleware

import (
	"net/http"
	"os"
	"signup-login-bumble/config"
	"signup-login-bumble/model"
	"signup-login-bumble/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func BearerAuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error": "Not in the login session, please login again"})
			return
		}

		data, err := utils.ValidateToken(token, os.Getenv("JWT_SECRET_KEY"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "error": err.Error()})
			return
		}

		var user model.User
		tokenUserData := data.(utils.JWTCustomClaims)
		result := config.DB.First(&user, "id = ?", tokenUserData.UserID)
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": result.Error.Error()})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
