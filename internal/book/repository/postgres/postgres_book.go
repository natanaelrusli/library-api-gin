package postgres

import (
	"database/sql"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type postgresBookRepository struct {
	Conn *sql.DB
}

func NewPostgresBookRepository(conn *sql.DB) domain.BookRepository {
	return &postgresBookRepository{
		Conn: conn,
	}
}

func (r *postgresBookRepository) FetchAll() (res []domain.Book, err error) {
	var books []domain.Book
	q := `
	SELECT * from books WHERE deleted_at IS NULL;
	`

	rows, err := r.Conn.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.Cover,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.DeletedAt,
			&book.AuthorID,
			&book.Stock,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
