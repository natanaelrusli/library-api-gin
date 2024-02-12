package domain

import "time"

type Book struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Cover       string    `json:"cover"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type BookUsecase interface {
}

type BookRepository interface {
}
