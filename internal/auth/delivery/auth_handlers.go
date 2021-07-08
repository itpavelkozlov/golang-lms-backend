package delivery

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandlers struct {
	AuthService domain.AuthService
	Logger      logger.Logger
}

func NewAuthHandlers(a domain.AuthService, logger logger.Logger) *AuthHandlers {
	return &AuthHandlers{
		AuthService: a,
		Logger:      logger,
	}
}

func (a *AuthHandlers) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello")
}
