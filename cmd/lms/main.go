package main

import (
	"context"
	"flag"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/database"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	var path = flag.String("config", "", "Path to config file")
	flag.Parse()
	if *path == "" {
		log.Fatalf("Please, provide --config")
	}

	c, err := config.NewConfig(*path)
	if err != nil {
		log.Fatalf(err.Error())
	}

	l, err := logger.NewLogger(c)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = database.NewDatabase(context.Background(), l, c)
	if err != nil {
		os.Exit(1)
	}

	l.Info("Database connected", zap.String("host", c.Service.Database.Host), zap.String("port", c.Service.Database.Port))

}
