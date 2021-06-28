package wire

import (
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Logger logger.Logger
	Echo   *echo.Echo
	Config *config.Config
}

func NewApplication(echo *echo.Echo, Config *config.Config, Logger logger.Logger) Application {
	return Application{
		Echo:   echo,
		Config: Config,
		Logger: Logger,
	}
}
