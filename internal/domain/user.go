package domain

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Phone     string       `json:"phone"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type UserUsecase interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchByName(ctx context.Context, name string) (User, error)
}

type UserRepository interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchByName(ctx context.Context, name string) (User, error)
}
