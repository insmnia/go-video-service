package routers

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-video-service/internal/service"
	"net/http"
	"strings"
)

type CategoryRouter struct {
	Service *service.CategoryService
}

func NewCategoryRouter(s *service.CategoryService) *CategoryRouter {
	return &CategoryRouter{s}
}

func (rt *CategoryRouter) NewCategoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			{
				rt.createCategoryRoute(w, r)
				return
			}
		case "GET":
			{
				categoryId := strings.TrimLeft(r.URL.String(), "/api/category/")
				if categoryId != "" {
					if _, err := uuid.Parse(categoryId); err != nil {
						http.Error(w, "Value is not a valid uuid", http.StatusUnprocessableEntity)
						return
					}
					rt.retrieveCategoryRoute(w, categoryId)
					return
				}
				rt.listCategoryRoute(w)
				return
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(`{"message": "Can't find method requested"}`))
			if err != nil {
				return
			}
		}
	}
}

func (rt *CategoryRouter) createCategoryRoute(w http.ResponseWriter, r *http.Request) {
	category, err := rt.Service.CreateCategory(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responseBytes, marshalErr := json.Marshal(category)
	if marshalErr != nil {
		panic(marshalErr)
	}
	_, writingErr := w.Write(responseBytes)
	if writingErr != nil {
		panic(writingErr)
	}
	w.WriteHeader(http.StatusOK)
}

func (rt *CategoryRouter) listCategoryRoute(w http.ResponseWriter) {
	categories := rt.Service.ListCategories()
	responseBytes, marshalErr := json.Marshal(categories)
	if marshalErr != nil {
		panic(marshalErr)
	}
	_, writingErr := w.Write(responseBytes)
	if writingErr != nil {
		panic(writingErr)
	}
	w.WriteHeader(http.StatusOK)
}
func (rt *CategoryRouter) retrieveCategoryRoute(w http.ResponseWriter, categoryId string) {
	category, err := rt.Service.GetCategory(categoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	responseBytes, marshalErr := json.Marshal(category)
	if marshalErr != nil {
		panic(marshalErr)
	}
	_, writingErr := w.Write(responseBytes)
	if writingErr != nil {
		panic(writingErr)
	}
	w.WriteHeader(http.StatusOK)
}
