package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-video-service/config"
	"time"
)

type JWTClaim struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
	jwt.StandardClaims
}

func ValidateToken(jwtConfig config.JWTConfig, signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtConfig.JwtSecret), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
