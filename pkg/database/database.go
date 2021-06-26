package database

import (
	"context"
	"fmt"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

func getDbUri(cfg config.Database) string {
	// example: postgres://username:password@localhost:5432/database_name
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}

func NewDatabase(ctx context.Context, logger logger.Logger, config *config.Config) (*pgx.Conn, error) {

	uri := getDbUri(config.Service.Database)

	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
		return nil, err
	}

	return conn, nil
}
