package main

import (
	"flag"
	"fmt"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"log"
	"os"
)

func main() {
	var path = flag.String("config", "", "Path to config file")
	flag.Parse()
	if *path == "" {
		os.Exit(1)
	}

	c, err := config.NewConfig(*path)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(c.Service.Name)
}
