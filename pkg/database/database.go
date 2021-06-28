package database

import (
	"fmt"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func getDbUri(cfg config.Database) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func NewDatabase(logger logger.Logger, config *config.Config) (*sqlx.DB, error) {

	conn, err := sqlx.Connect("pgx", getDbUri(config.Service.Database))
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
		return nil, err
	}
	logger.Info("Database connected", zap.String("host", config.Service.Database.Host), zap.String("port", config.Service.Database.Port))
	return conn, nil
}
