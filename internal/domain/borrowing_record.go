package domain

import (
	"context"
	"database/sql"
	"time"
)

type BorrowingRecord struct {
	Id            int          `json:"id"`
	UserId        int          `json:"user_id"`
	BookId        int          `json:"book_id"`
	Status        string       `json:"status"`
	BorrowingDate time.Time    `json:"borrowing_date"`
	ReturningDate sql.NullTime `json:"returning_date"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
}

type BorrowingRecordUsecase interface {
	CreateRecord(ctx context.Context, userId int, bookId int, status string) (BorrowingRecord, error)
	GetAllBorrowedRecord(ctx context.Context) ([]BorrowingRecord, error)
}

type BorrowingRecordRepository interface {
	CreateRecord(ctx context.Context, record BorrowingRecord) (BorrowingRecord, error)
	GetAllBorrowedRecord(ctx context.Context) ([]BorrowingRecord, error)
}
