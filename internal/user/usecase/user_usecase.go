package usecase

import (
	"context"

	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) FetchAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.userRepo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) FetchByName(ctx context.Context, name string) (domain.User, error) {
	user, err := u.userRepo.FetchByName(ctx, name)
	if err != nil {
		return domain.User{}, nil
	}

	return user, nil
}
