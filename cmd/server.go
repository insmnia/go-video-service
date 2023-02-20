package cmd

import (
	"go-video-service/config"
	"go-video-service/database"
	"go-video-service/internal/api"
	"go-video-service/internal/middlewares"
	"log"
	"net/http"
)

func InitServer() *http.ServeMux {
	dbConf, dbConfErr := config.LoadDatabaseConfig("env/")
	if dbConfErr != nil {
		log.Fatalln(dbConfErr)
	}
	db, dbErr := database.InitDB(dbConf)
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	jwtConf, jwtConfErr := config.LoadJWTConfig("env/")
	if jwtConfErr != nil {
		log.Fatalln(jwtConfErr)
	}
	mux := http.NewServeMux()

	mux.Handle("/api/category/", middlewares.AuthMiddleware(api.NewCategoryAPIHandler(db), jwtConf))
	return mux
}
