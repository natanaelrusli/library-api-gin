package postgres

import (
	"context"
	"database/sql"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
)

type postgresBookRepository struct {
	Conn *sql.DB
}

func NewPostgresBookRepository(conn *sql.DB) domain.BookRepository {
	return &postgresBookRepository{
		Conn: conn,
	}
}

func (r *postgresBookRepository) FetchAll(ctx context.Context) (res []domain.Book, err error) {
	var books []domain.Book
	q := `
		SELECT * from books WHERE deleted_at IS NULL;
	`

	rows, err := r.Conn.QueryContext(ctx, q)
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

func (r *postgresBookRepository) GetByID(ctx context.Context, id int) (domain.Book, error) {
	var book domain.Book
	q := `
		SELECT * from books WHERE id = $1 AND deleted_at IS NULL;
	`

	err := r.Conn.QueryRowContext(ctx, q, id).Scan(
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

func (r *postgresBookRepository) CreateOne(ctx context.Context, book domain.Book) (domain.Book, error) {
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

	result := r.Conn.QueryRowContext(ctx, q,
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

func (r *postgresBookRepository) FetchAllWithAuthor(ctx context.Context) ([]domain.BookWithAuthor, error) {
	var booksWithAuthor []domain.BookWithAuthor
	q := `
		SELECT 
			b.id,
			b.title,
			b.description,
			b.cover,
			b.author_id,
			a.name author_name,
			b.stock,
			b.updated_at,
			b.created_at,
			b.deleted_at 
		FROM books AS b
		JOIN authors as a ON b.author_id = a.id 
		WHERE b.deleted_at IS NULL;
	`

	rows, err := r.Conn.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookWithAuthor domain.BookWithAuthor
		if err := rows.Scan(
			&bookWithAuthor.Id,
			&bookWithAuthor.Title,
			&bookWithAuthor.Description,
			&bookWithAuthor.Cover,
			&bookWithAuthor.AuthorID,
			&bookWithAuthor.AuthorName,
			&bookWithAuthor.Stock,
			&bookWithAuthor.UpdatedAt,
			&bookWithAuthor.CreatedAt,
			&bookWithAuthor.DeletedAt,
		); err != nil {
			return nil, err
		}
		booksWithAuthor = append(booksWithAuthor, bookWithAuthor)
	}

	return booksWithAuthor, nil
}

func (r *postgresBookRepository) UpdateStock(ctx context.Context, req dto.BorrowRequest) (domain.Book, error) {
	var book domain.Book

	q := `
		UPDATE books
		SET stock = $1
		WHERE id = $2
		RETURNING *;
	`
	err := r.Conn.QueryRowContext(ctx, q, req.Amount, req.BookId).Scan(
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
