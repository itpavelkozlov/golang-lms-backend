// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/server"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/delivery/http"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/repository/postgres"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/usecase"
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
	userRepository := postgres.NewPostgresUserRepository(db, loggerLogger)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := http.NewUserHandler(userUsecase, loggerLogger)
	echo := server.NewHttpServer(userHandler)
	application := NewApplication(echo, configConfig, loggerLogger)
	return application, nil
}
