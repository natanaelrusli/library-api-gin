package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/pkg/apperror"
)

type bookUsecase struct {
	bookRepo   domain.BookRepository
	authorRepo domain.AuthorRepository
}

func NewBookUsecase(br domain.BookRepository, ar domain.AuthorRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepo:   br,
		authorRepo: ar,
	}
}

func (u *bookUsecase) FetchAll(ctx context.Context) ([]domain.Book, error) {
	books, err := u.bookRepo.FetchAll(ctx)

	if err != nil && err == sql.ErrNoRows {
		return nil, apperror.NewNotFoundError(nil, "book")
	}

	return books, nil
}

func (u *bookUsecase) GetByID(ctx context.Context, id int) (domain.Book, error) {
	book, err := u.bookRepo.GetByID(ctx, id)

	if err != nil {
		return domain.Book{}, apperror.NewBookNotFoundError()
	}

	return book, nil
}

func (u *bookUsecase) CreateOne(
	ctx context.Context,
	title string,
	description string,
	cover string,
	authorId int32,
	stock int32,
) (domain.Book, error) {
	var book domain.Book

	book.Title = title
	book.Description = description
	book.Cover = cover
	book.AuthorID = authorId
	book.Stock = stock
	book.UpdatedAt = time.Now()
	book.CreatedAt = time.Now()

	resultBook, err := u.bookRepo.CreateOne(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}

	return resultBook, nil
}

func (u *bookUsecase) GetBookAuthor(ctx context.Context, id int) (domain.Author, error) {
	book, err := u.bookRepo.GetByID(ctx, id)

	if err != nil {
		return domain.Author{}, err
	}

	author, err := u.authorRepo.GetByID(ctx, int64(book.AuthorID))

	if err != nil {
		return domain.Author{}, err
	}

	return author, nil
}

func (u *bookUsecase) FetchAllWithAuthor(ctx context.Context) ([]domain.BookWithAuthor, error) {
	books, err := u.bookRepo.FetchAllWithAuthor(ctx)

	if err != nil {
		return []domain.BookWithAuthor{}, err
	}

	return books, nil
}
