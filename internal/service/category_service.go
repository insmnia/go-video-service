package service

import (
	"database/sql"
	"encoding/json"
	"go-video-service/internal/api/dto"
	"go-video-service/internal/repository"
	"go-video-service/pkg/errors"
	"io"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) CreateCategory(categoryUpsert io.ReadCloser) (dto.CategoryResponse, error) {
	decoder := json.NewDecoder(categoryUpsert)
	var input dto.CreateCategoryRequest
	err := decoder.Decode(&input)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	_, _, isExist := s.Repo.GetByName(input.Name)
	if isExist {
		return dto.CategoryResponse{}, &errors.CategoryAlreadyExists{}
	}
	category := s.Repo.Create(input.Name)
	return dto.CategoryResponse{
		Id:        category.Id,
		Name:      category.CategoryName,
		CreatedAt: category.CreatedAt,
		DeletedAt: category.DeletedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (s *CategoryService) ListCategories() []dto.CategoryResponse {
	categories := s.Repo.List()
	dtoListResponse := make([]dto.CategoryResponse, len(categories))
	for _, c := range categories {
		dtoListResponse = append(dtoListResponse, dto.CategoryResponse{
			Id:        c.Id,
			Name:      c.CategoryName,
			CreatedAt: c.CreatedAt,
			DeletedAt: c.DeletedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return dtoListResponse
}

func (s *CategoryService) GetCategory(categoryId string) (dto.CategoryResponse, error) {
	category, err := s.Repo.GetById(categoryId)
	switch err {
	case nil:
		{
			return dto.CategoryResponse{
				Id:        category.Id,
				Name:      category.CategoryName,
				CreatedAt: category.CreatedAt,
				DeletedAt: category.DeletedAt,
				UpdatedAt: category.UpdatedAt,
			}, nil
		}
	case sql.ErrNoRows:
		{
			return dto.CategoryResponse{}, &errors.CategoryNotFound{}
		}
	default:
		panic(err)
	}
}
