package usecase

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func (u userUsecase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	if len(id) != 36 {
		return nil, domain.ErrBadParamInput
	}
	return u.userRepo.GetByID(ctx, id)
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}
