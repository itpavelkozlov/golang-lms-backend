package server

import (
	auth "github.com/itpavelkozlov/golang-lms-backend/internal/auth/delivery"
	user "github.com/itpavelkozlov/golang-lms-backend/internal/user/delivery"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func NewHttpServer(userHandler *user.UserHandlers, authHandlers *auth.AuthHandlers) *echo.Echo {
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.GET("/users/:id", userHandler.GetByID)
	e.GET("/login", authHandlers.Login)
	return e
}
