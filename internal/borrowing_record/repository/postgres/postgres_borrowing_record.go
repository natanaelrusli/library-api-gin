package postgres

import (
	"context"
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

func (r *postgresBorrowingRecordRepository) CreateRecord(ctx context.Context, record domain.BorrowingRecord) (domain.BorrowingRecord, error) {
	var createdRecord domain.BorrowingRecord

	q := `
		INSERT INTO borrowing_records
		(user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING
		id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at;
	`

	err := r.Conn.QueryRowContext(ctx, q,
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

func (r *postgresBorrowingRecordRepository) GetAllBorrowedRecord(ctx context.Context) ([]domain.BorrowingRecord, error) {
	var records []domain.BorrowingRecord
	q := `
		SELECT * 
		FROM borrowing_records
		WHERE status = 'BORROWED' 
		AND deleted_at IS NULL;	
	`

	rows, err := r.Conn.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var record domain.BorrowingRecord

		if err := rows.Scan(
			&record.Id,
			&record.UserId,
			&record.BookId,
			&record.Status,
			&record.BorrowingDate,
			&record.ReturningDate,
			&record.CreatedAt,
			&record.UpdatedAt,
			&record.DeletedAt,
		); err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}
