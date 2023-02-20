package api

import (
	"github.com/jmoiron/sqlx"
	"go-video-service/internal/api/routers"
	"go-video-service/internal/repository"
	"go-video-service/internal/service"
	"net/http"
)

func NewCategoryAPIHandler(db *sqlx.DB) http.HandlerFunc {
	repo := repository.NewCategoryRepository(db)
	s := service.NewCategoryService(repo)
	router := routers.NewCategoryRouter(s)
	return router.NewCategoryHandler()
}
