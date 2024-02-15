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
	var createdRecord domain.BorrowingRecord

	q := `
		INSERT INTO borrowing_records 
		(user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING
		id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at;
	`

	err := r.Conn.QueryRow(q,
		&record.UserId,
		&record.BookId,
		&record.Status,
		&record.BorrowingDate,
		&record.ReturningDate,
		&record.CreatedAt,
		&record.UpdatedAt,
		&record.DeletedAt,
	).Scan(
		&createdRecord.Id,
		&createdRecord.UserId,
		&createdRecord.BookId,
		&createdRecord.Status,
		&createdRecord.BorrowingDate,
		&createdRecord.ReturningDate,
		&createdRecord.CreatedAt,
		&createdRecord.UpdatedAt,
		&createdRecord.DeletedAt,
	)

	if err != nil {
		return domain.BorrowingRecord{}, err
	}

	return createdRecord, nil
}
