package server

import (
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/delivery"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func NewHttpServer(userHandler *delivery.UserHandlers) *echo.Echo {
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.GET("/users/:id", userHandler.GetByID)
	return e
}
