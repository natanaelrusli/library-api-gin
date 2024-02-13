package usecase

import "github.com/natanaelrusli/library-api-gin/internal/domain"

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
