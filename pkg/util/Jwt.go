package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/trananh-it-hust/ChatApp/global"
)

type Claims struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userID int, email string) (string, error) {
	conf := global.Config
	secretKey := []byte(conf.Jwt.Secret)
	expTime := conf.Jwt.Expire
	expirationTime := time.Now().Add(time.Duration(expTime) * time.Minute)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	conf := global.Config
	secretKey := []byte(conf.Jwt.Secret)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
