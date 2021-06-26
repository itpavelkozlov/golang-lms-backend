package server

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/delivery/http"
	"github.com/labstack/echo/v4"
)

func NewHttpServer(userHandler *http.UserHandler) *echo.Echo {
	e := echo.New()
	e.GET("/users/:id", userHandler.GetByID)
	return e
}
