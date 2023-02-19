package middlewares

import (
	"go-video-service/config"
	"go-video-service/pkg/utils"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, jwtConf config.JWTConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Couldn't parse token", http.StatusBadRequest)
			return
		}
		tokenParts := strings.Split(tokenString, " ")
		claims, err := utils.ValidateToken(jwtConf, tokenParts[1])
		if err != nil {
			log.Print(err)
			http.Error(w, "Couldn't parse user claims", http.StatusUnauthorized)
			return
		}
		r.Header.Set("userId", claims.UserId)
		next.ServeHTTP(w, r)
	})
}
