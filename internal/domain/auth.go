package domain

import (
	"context"
	"time"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	AccessToken  string     `json:"access_token" db:"-"`
	RefreshToken string     `json:"refresh_token" db:"refresh_token"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

type AuthUsecase interface {
	GetNewClaims(ctx context.Context, auth *AuthRequest) (*Claims, error)
	RefreshClaims(ctx context.Context, claims *Claims) (*Claims, error)
	Logout(ctx context.Context, claims *Claims) error
}

type AuthRepository interface {
	UpdateSession(ctx context.Context, userId string) (*Claims, error)
	CreateSession(ctx context.Context, userId string) (*Claims, error)
	DeleteSession(ctx context.Context, userId string) error
	GetSessionByUserID(ctx context.Context, userId string) (string, error)
}
