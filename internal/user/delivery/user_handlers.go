package delivery

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/internal/helpers"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandlers struct {
	UserService domain.UserService
	Logger      logger.Logger
}

func NewUserHandlers(u domain.UserService, logger logger.Logger) *UserHandlers {
	return &UserHandlers{
		UserService: u,
		Logger:      logger,
	}
}

func (a *UserHandlers) GetByID(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()

	u, err := a.UserService.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}
