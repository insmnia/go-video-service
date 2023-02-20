package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go-video-service/models"
	"log"
)

type CategoryRepository struct {
	Db *sqlx.DB
}

func NewCategoryRepository(Db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{Db}
}

func (repo *CategoryRepository) Create(categoryName string) models.CategoryModel {
	var newCategory models.CategoryModel
	err := repo.Db.QueryRowx(
		`INSERT INTO category (category_name) VALUES ($1) RETURNING *;`, categoryName,
	).StructScan(&newCategory)
	if err != nil {
		log.Fatal(err)
	}
	return newCategory
}

func (repo *CategoryRepository) GetByName(categoryName string) (category models.CategoryModel, err error, found bool) {
	err = repo.Db.Get(&category, "SELECT * FROM category WHERE category_name=$1", categoryName)
	switch err {
	case sql.ErrNoRows:
		{
			return category, err, false
		}
	case nil:
		{
			return category, nil, true

		}
	default:
		log.Fatalln(err)
	}
	return
}

func (repo *CategoryRepository) List() (categories []models.CategoryModel) {
	err := repo.Db.Select(&categories, "SELECT * FROM category WHERE deleted_at is NULL")
	if err != nil {
		panic(err)
	}
	return
}
func (repo *CategoryRepository) GetById(categoryId string) (category models.CategoryModel, err error) {
	err = repo.Db.Get(&category, "SELECT * FROM category WHERE id=$1", categoryId)
	return
}
