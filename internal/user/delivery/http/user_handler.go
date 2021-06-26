package http

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

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
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}

// Todo check on save
func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
