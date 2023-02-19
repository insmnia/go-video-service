package cmd

import (
	"encoding/json"
	"go-video-service/config"
	"go-video-service/database"
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

	indexHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec("SELECT * FROM category")
		if err != nil {
			return
		}
		data, err := json.Marshal(r.Header.Get("userId"))
		if err != nil {
			log.Fatalln(err)
		}
		_, err = w.Write(data)
		if err != nil {
			return
		}
	})
	mux.Handle("/", middlewares.AuthMiddleware(indexHandler, jwtConf))
	return mux
}
