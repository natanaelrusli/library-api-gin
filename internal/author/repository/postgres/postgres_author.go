package postgres

import (
	"context"
	"database/sql"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type postgresAuthorRepository struct {
	Conn *sql.DB
}

func NewPostgresAuthorRepository(conn *sql.DB) domain.AuthorRepository {
	return &postgresAuthorRepository{
		Conn: conn,
	}
}

func (r *postgresAuthorRepository) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	var author domain.Author
	q := `
		SELECT * from authors
		WHERE deleted_at IS NULL
		AND id = $1;
	`

	rows := r.Conn.QueryRowContext(ctx, q, id)
	if rows.Err() != nil {
		return domain.Author{}, rows.Err()
	}

	err := rows.Scan(
		&author.ID,
		&author.Name,
		&author.CreatedAt,
		&author.UpdatedAt,
		&author.DeletedAt,
	)

	if err != nil {
		return domain.Author{}, err
	}

	return author, nil
}
