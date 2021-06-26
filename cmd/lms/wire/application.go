package wire

import (
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Echo   *echo.Echo
	Config *config.Config
}

func NewApplication(echo *echo.Echo, Config *config.Config) Application {
	return Application{
		Echo:   echo,
		Config: Config,
	}
}
