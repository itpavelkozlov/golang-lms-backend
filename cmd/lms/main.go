package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/itpavelkozlov/golang-lms-backend/cmd/lms/wire"
	"log"
)

func main() {
	var path = flag.String("config", "", "Path to config file")
	flag.Parse()
	if *path == "" {
		log.Fatalf("Please, provide --config")
	}

	ctx := context.Background()
	_, err := wire.InitializeApp(ctx, *path)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error when app initializing: %s", err.Error()))
	}
}
