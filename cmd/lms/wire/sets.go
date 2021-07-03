package wire

import (
	"github.com/google/wire"
	"github.com/itpavelkozlov/golang-lms-backend/internal/server"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/delivery"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/repository"
	"github.com/itpavelkozlov/golang-lms-backend/internal/user/service"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/database"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
)

var pkgSet = wire.NewSet(
	database.NewDatabase,
	logger.NewLogger,
	config.NewConfig,
)

var userSet = wire.NewSet(
	service.NewUserService,
	repository.NewPostgresUserRepository,
	delivery.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHttpServer,
)
