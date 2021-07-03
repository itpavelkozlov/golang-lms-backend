package wire

import (
	"github.com/google/wire"
	authDelivery "github.com/itpavelkozlov/golang-lms-backend/internal/auth/delivery"
	authRepository "github.com/itpavelkozlov/golang-lms-backend/internal/auth/repository"
	authService "github.com/itpavelkozlov/golang-lms-backend/internal/auth/service"
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
	repository.NewUserRepository,
	delivery.NewUserHandlers,
)

var authSet = wire.NewSet(
	authService.NewAuthService,
	authRepository.NewAuthRepository,
	authDelivery.NewAuthHandlers,
)

var serverSet = wire.NewSet(
	server.NewHttpServer,
)
