package models

import (
	"github.com/google/uuid"
	"time"
)

type CategoryModel struct {
	Id           uuid.UUID
	CategoryName string     `db:"category_name"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
