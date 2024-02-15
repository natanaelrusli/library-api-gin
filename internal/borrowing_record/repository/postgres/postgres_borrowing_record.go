package postgres

import (
	"database/sql"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type postgresBorrowingRecordRepository struct {
	Conn *sql.DB
}

func NewBorrowingRecordRepository(conn *sql.DB) domain.BorrowingRecordRepository {
	return &postgresBorrowingRecordRepository{
		Conn: conn,
	}
}

func (r *postgresBorrowingRecordRepository) CreateRecord(record domain.BorrowingRecord) (domain.BorrowingRecord, error) {
	q := `
		INSERT INTO borrowing_records 
		(id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING
		id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at;
	`

	r.Conn.QueryRow(q)

	return domain.BorrowingRecord{}, nil
}
