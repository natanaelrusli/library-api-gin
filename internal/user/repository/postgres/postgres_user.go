package postgres

import (
	"database/sql"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

func NewPostgresUserRepository(conn *sql.DB) domain.UserRepository {
	return &postgresUserRepository{
		Conn: conn,
	}
}

func (r *postgresUserRepository) FetchAll() ([]domain.User, error) {
	var users []domain.User
	q := `
		SELECT * 
		FROM users 
		WHERE deleted_at IS NULL;
	`

	rows, err := r.Conn.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.Email,
			&user.Password,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
