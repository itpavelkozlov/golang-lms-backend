package service

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
)

type userService struct {
	userRepository domain.UserRepository
}

func (u userService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	if len(id) != 36 {
		return nil, domain.ErrBadParamInput
	}
	return u.userRepository.GetByID(ctx, id)
}

func NewUserService(u domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: u,
	}
}
