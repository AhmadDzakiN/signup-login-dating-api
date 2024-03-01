package utils

import (
	"os"
	"signup-login-bumble/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID   string `json:"id"`
	Fullname string `json:"full_name"`
}

func CreateToken(userData model.User) (token string, err error) {
	customClaims := &JWTCustomClaims{
		UserID:   userData.ID,
		Fullname: userData.Fullname,
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub": customClaims,
		"exp": now.Unix(),
		"nbf": now.Unix(),
		"iat": now.Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return
	}

	return
}
