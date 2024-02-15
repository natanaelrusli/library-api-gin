package domain

import (
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
	CreateRecord(userId int, bookId int, status string) (BorrowingRecord, error)
	GetAllBorrowedRecord() ([]BorrowingRecord, error)
}

type BorrowingRecordRepository interface {
	CreateRecord(record BorrowingRecord) (BorrowingRecord, error)
	GetAllBorrowedRecord() ([]BorrowingRecord, error)
}
