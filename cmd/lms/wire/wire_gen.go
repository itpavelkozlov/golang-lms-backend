// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/server"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/delivery"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/repository"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/service"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/database"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, configPath string) (Application, error) {
	configConfig, err := config.NewConfig(configPath)
	if err != nil {
		return Application{}, err
	}
	loggerLogger, err := logger.NewLogger(configConfig)
	if err != nil {
		return Application{}, err
	}
	db, err := database.NewDatabase(loggerLogger, configConfig)
	if err != nil {
		return Application{}, err
	}
	userRepository := repository.NewPostgresUserRepository(db, loggerLogger)
	userService := service.NewUserService(userRepository)
	userHandlers := delivery.NewUserHandler(userService, loggerLogger)
	echo := server.NewHttpServer(userHandlers)
	application := NewApplication(echo, configConfig, loggerLogger)
	return application, nil
}
