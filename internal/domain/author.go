package domain

import (
	"context"
	"database/sql"
	"time"
)

type Author struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}
