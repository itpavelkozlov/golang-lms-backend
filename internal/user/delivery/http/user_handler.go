package http

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/internal/helpers"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
	Logger      logger.Logger
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(u domain.UserUsecase, logger logger.Logger) *UserHandler {
	return &UserHandler{
		UserUsecase: u,
		Logger:      logger,
	}
}

// GetByID will get user by given id
func (a *UserHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()

	u, err := a.UserUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}

// Login will get accessToken by given username/password
func (a *UserHandler) Login(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()

	u, err := a.UserUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}
