package domain

import (
	"context"
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

type BookWithAuthor struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Cover       string       `json:"cover"`
	AuthorID    int32        `json:"author_id"`
	AuthorName  string       `json:"author_name"`
	Stock       int32        `json:"stock"`
	UpdatedAt   time.Time    `json:"updated_at"`
	CreatedAt   time.Time    `json:"created_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

type BookUsecase interface {
	FetchAll(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int) (Book, error)
	CreateOne(
		ctx context.Context,
		title string,
		description string,
		cover string,
		authorId int32,
		stock int32,
	) (Book, error)
	GetBookAuthor(ctx context.Context, id int) (Author, error)
	FetchAllWithAuthor(ctx context.Context) ([]BookWithAuthor, error)
}

type BookRepository interface {
	FetchAll(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int) (Book, error)
	CreateOne(ctx context.Context, book Book) (Book, error)
	FetchAllWithAuthor(ctx context.Context) ([]BookWithAuthor, error)
	UpdateStock(ctx context.Context, stock int, bookId int) (Book, error)
}
