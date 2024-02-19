package usecase

import (
	"context"
	"log"
	"testing"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBookUsecase(t *testing.T) {
	dep := NewBookUsecase(nil, nil, nil)
	assert.NotNil(t, dep)
}

func TestFetchAll(t *testing.T) {
	type fields struct {
		bookRepository *mocks.BookRepository
	}

	tests := []struct {
		name       string
		wantResult []domain.Book
		mock       func(ctx context.Context, f fields)
	}{
		{
			name: "success",
			mock: func(ctx context.Context, f fields) {
				f.bookRepository.On(
					"FetchAll",
					ctx,
					mock.Anything,
				).Return(
					[]domain.Book{
						{
							Id:          1,
							Title:       "title",
							Description: "description",
						},
					},
				)
			},
		},
	}

	for _, tt := range tests {
		log.Println(tt.name)
	}
}
