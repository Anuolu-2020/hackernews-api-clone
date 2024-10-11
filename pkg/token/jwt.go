package token

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Anuolu-2020/hackernews-api-clone/pkg/env"
)

var (
	secretKey string
	t         *jwt.Token
)

func GenerateToken(username string) (string, error) {
	secretKey = pkg.GetEnv("JWT_SECRET_TOKEN")

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := t.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalf("Error while generating token: %v", err)
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	secretKey := pkg.GetEnv("JWT_SECRET_TOKEN")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
