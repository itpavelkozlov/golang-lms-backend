package delivery

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
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

func (a *AuthHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello")
}
