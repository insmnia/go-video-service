package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-video-service/config"
	"log"
)

func InitDB(config config.DatabaseConfig) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
			config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
