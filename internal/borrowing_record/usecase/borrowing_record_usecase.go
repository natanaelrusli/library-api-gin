package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/natanaelrusli/library-api-gin/internal/constants"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
)

type borrowingRecordUsecase struct {
	borrowingRecordRepo domain.BorrowingRecordRepository
	bookRepo            domain.BookRepository
}

func NewBorrowingRecordUsecase(brr domain.BorrowingRecordRepository, br domain.BookRepository) domain.BorrowingRecordUsecase {
	return &borrowingRecordUsecase{
		borrowingRecordRepo: brr,
		bookRepo:            br,
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

func (u *borrowingRecordUsecase) Borrow(ctx context.Context, id int, amount int) (domain.BorrowingRecord, error) {
	// check stock availability
	book, err := u.bookRepo.GetByID(ctx, id)
	userId := 1

	if err != nil {
		return domain.BorrowingRecord{}, err
	}

	if book.Stock < int32(amount) {
		return domain.BorrowingRecord{}, errors.New("book stock not enough")
	}

	newAmount := book.Stock - int32(amount)

	req := dto.UpdateBookStockRequest{
		BookId: id,
		Amount: int(newAmount),
	}
	book, err = u.bookRepo.UpdateStock(ctx, req)

	if err != nil {
		return domain.BorrowingRecord{}, nil
	}

	// create borrowing record
	record := domain.BorrowingRecord{
		UserId:        userId,
		BookId:        id,
		Status:        constants.Borrowed,
		BorrowingDate: time.Now(),
		ReturningDate: sql.NullTime{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
	}

	record, err = u.borrowingRecordRepo.CreateRecord(ctx, record)
	if err != nil {
		return domain.BorrowingRecord{}, err
	}

	return record, nil
}
