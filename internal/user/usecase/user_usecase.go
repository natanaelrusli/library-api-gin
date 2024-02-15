package usecase

import "github.com/natanaelrusli/library-api-gin/internal/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) FetchAll() ([]domain.User, error) {
	users, err := u.userRepo.FetchAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) FetchByName(name string) (domain.User, error) {
	user, err := u.userRepo.FetchByName(name)
	if err != nil {
		return domain.User{}, nil
	}

	return user, nil
}
