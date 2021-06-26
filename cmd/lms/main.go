package main

import (
	"flag"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"go.uber.org/zap"
	"log"
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
	l.Info("Hello World", zap.String("Name", c.Service.Name))

}
