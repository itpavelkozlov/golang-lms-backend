package http

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
	Logger      logger.Logger
}

func NewAuthHandler(a domain.AuthUsecase, logger logger.Logger) *AuthHandler {
	return &AuthHandler{
		AuthUsecase: a,
		Logger:      logger,
	}
}
