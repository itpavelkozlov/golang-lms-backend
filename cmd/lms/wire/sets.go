package wire

import (
	"github.com/google/wire"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/database"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
)

var pkgSet = wire.NewSet(
	database.NewDatabase,
	logger.NewLogger,
	config.NewConfig,
)
