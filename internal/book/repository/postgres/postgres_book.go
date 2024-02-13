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

func (r *postgresBookRepository) GetByID(id int) (domain.Book, error) {
	var book domain.Book
	q := `
		SELECT * from books WHERE id = $1 AND deleted_at IS NULL;
	`

	err := r.Conn.QueryRow(q, id).Scan(
		&book.Id,
		&book.Title,
		&book.Description,
		&book.Cover,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.DeletedAt,
		&book.AuthorID,
		&book.Stock,
	)
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (r *postgresBookRepository) CreateOne(book domain.Book) (domain.Book, error) {
	q := `
		INSERT INTO books (
			title,
			description,
			cover,
			author_id,
			stock,
			updated_at,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, title, description, cover, author_id, stock, updated_at, created_at, deleted_at;
	`

	result := r.Conn.QueryRow(q,
		book.Title,
		book.Description,
		book.Cover,
		book.AuthorID,
		book.Stock,
		book.UpdatedAt,
		book.CreatedAt,
	)

	var resultBook domain.Book
	err := result.Scan(
		&resultBook.Id,
		&resultBook.Title,
		&resultBook.Description,
		&resultBook.Cover,
		&resultBook.AuthorID,
		&resultBook.Stock,
		&resultBook.UpdatedAt,
		&resultBook.CreatedAt,
		&resultBook.DeletedAt,
	)

	if err != nil {
		return domain.Book{}, err
	}

	return resultBook, nil
}
