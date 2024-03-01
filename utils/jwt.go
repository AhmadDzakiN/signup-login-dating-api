package utils

import (
	"errors"
	"fmt"
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
		"exp": now.Add(time.Minute * 43200).Unix(),
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

func ValidateToken(token, secretKey string) (data interface{}, err error) {

	extractedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in token parsing")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		err = fmt.Errorf("invalidate token: %w", err)
		return
	}

	if extractedToken == nil {
		err = errors.New("invalid token")
		return
	}

	claims, ok := extractedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("token error")
		return
	}

	claimsSubject := claims["sub"].(map[string]interface{})
	data = JWTCustomClaims{
		UserID:   claimsSubject["id"].(string),
		Fullname: claimsSubject["full_name"].(string),
	}

	return
}
