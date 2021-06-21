package main

import (
	"fmt"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err := config.NewConfig("./configs/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Hello LMS")
	fmt.Println(viper.Get("service.name"))
}
