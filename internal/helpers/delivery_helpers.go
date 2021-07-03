package helpers

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

func IsRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetStatusCode(err error) int {
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
