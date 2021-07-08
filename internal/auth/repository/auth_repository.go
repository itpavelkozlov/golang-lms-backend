package repository

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func (a authRepository) UpdateSession(ctx context.Context, userId string) (*domain.Claims, error) {
	panic("implement me")
}

func (a authRepository) CreateSession(ctx context.Context, userId string) (*domain.Claims, error) {
	panic("implement me")
}

func (a authRepository) DeleteSession(ctx context.Context, userId string) error {
	panic("implement me")
}

func (a authRepository) GetSessionByUserID(ctx context.Context, userId string) (string, error) {
	panic("implement me")
}

func NewAuthRepository(conn *sqlx.DB, logger logger.Logger) domain.AuthRepository {
	return &authRepository{conn, logger}
}
