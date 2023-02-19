package cmd

import (
	"encoding/json"
	"go-video-service/config"
	"go-video-service/database"
	"log"
	"net/http"
)

func InitServer() {
	conf, confErr := config.LoadDatabaseConfig("env/")
	if confErr != nil {
		log.Fatalln(confErr)
	}
	db, dbErr := database.InitDB(conf)
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec("SELECT * FROM category")
		if err != nil {
			return
		}
		data, err := json.Marshal("Hi")
		if err != nil {
			log.Fatalln(err)
		}
		_, err = w.Write(data)
		if err != nil {
			return
		}
	})
}
