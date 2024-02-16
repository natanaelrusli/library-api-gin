package usecase

import (
	"context"
	"time"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
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

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) GetByID(ctx context.Context, id int) (domain.Book, error) {
	book, err := u.bookRepo.GetByID(ctx, id)

	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (u *bookUsecase) CreateOne(
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

	resultBook, err := u.bookRepo.CreateOne(book)
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

	author, err := u.authorRepo.GetByID(int64(book.AuthorID))

	if err != nil {
		return domain.Author{}, err
	}

	return author, nil
}

func (u *bookUsecase) FetchAllWithAuthor() ([]domain.BookWithAuthor, error) {
	books, err := u.bookRepo.FetchAllWithAuthor()

	if err != nil {
		return []domain.BookWithAuthor{}, err
	}

	return books, nil
}
