package usecase

import (
	"time"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type bookUsecase struct {
	bookRepo domain.BookRepository
}

func NewBookUsecase(br domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepo: br,
	}
}

func (u *bookUsecase) FetchAll() ([]domain.Book, error) {
	books, err := u.bookRepo.FetchAll()

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) GetByID(id int) (domain.Book, error) {
	book, err := u.bookRepo.GetByID(id)

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
