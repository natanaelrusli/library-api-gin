package usecase

import (
	"context"
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

func (u *borrowingRecordUsecase) CreateRecord(ctx context.Context, userId int, bookId int, status string) (domain.BorrowingRecord, error) {
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

	record, err := u.borrowingRecordRepo.CreateRecord(ctx, record)
	if err != nil {
		return domain.BorrowingRecord{}, err
	}

	return record, nil
}

func (u *borrowingRecordUsecase) GetAllBorrowedRecord(ctx context.Context) ([]domain.BorrowingRecord, error) {
	records, err := u.borrowingRecordRepo.GetAllBorrowedRecord(ctx)
	if err != nil {
		return nil, err
	}

	return records, nil
}
