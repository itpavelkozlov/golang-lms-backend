package service

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
)

type authService struct {
	authRepository domain.AuthRepository
}

func (a authService) GetNewClaims(ctx context.Context, auth *domain.AuthRequest) (*domain.Claims, error) {
	panic("implement me")
}

func (a authService) RefreshClaims(ctx context.Context, claims *domain.Claims) (*domain.Claims, error) {
	panic("implement me")
}

func (a authService) Logout(ctx context.Context, claims *domain.Claims) error {
	panic("implement me")
}

func NewAuthService(a domain.AuthRepository) domain.AuthService {
	return &authService{
		authRepository: a,
	}
}
