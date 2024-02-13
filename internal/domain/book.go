package domain

import (
	"database/sql"
	"time"
)

type Book struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Cover       string       `json:"cover"`
	AuthorID    int32        `json:"author_id"`
	Stock       int32        `json:"stock"`
	UpdatedAt   time.Time    `json:"updated_at"`
	CreatedAt   time.Time    `json:"created_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

type BookUsecase interface {
	FetchAll() ([]Book, error)
	GetByID(id int) (Book, error)
}

type BookRepository interface {
	FetchAll() ([]Book, error)
	GetByID(id int) (Book, error)
}
