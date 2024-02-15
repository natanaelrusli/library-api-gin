package usecase

import (
	"database/sql"
	"time"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type borrowingRecordUsecase struct {
	borrowingRecordRepo domain.BorrowingRecordRepository
}

func NewBorrowingRecordUsecase(brr domain.BorrowingRecordRepository) domain.BorrowingRecordUsecase {
	return &borrowingRecordUsecase{
		borrowingRecordRepo: brr,
	}
}

func (u *borrowingRecordUsecase) CreateRecord(userId int, bookId int, status string) (domain.BorrowingRecord, error) {
	record := domain.BorrowingRecord{
		UserId:        userId,
		BookId:        bookId,
		Status:        status,
		BorrowingDate: time.Now(),
		ReturningDate: sql.NullTime{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
	}

	record, err := u.borrowingRecordRepo.CreateRecord(record)
	if err != nil {
		return domain.BorrowingRecord{}, err
	}

	return record, nil
}
