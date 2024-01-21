package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"subscriber/internal/handler"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
	"subscriber/pkg/server"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	fmt.Println("server started")
	if err := methods.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	nc, err := methods.InitNatsBridge()
	if err != nil {
		log.Fatalln("error initializing nats bridge: %s", err.Error())
	}

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.PersonData{})

	handlers := handler.NewHandler()

	subscriber := methods.SubscribeModelOrder(db, nc)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("app_port"), handlers.InitAuthRoutes()); err != nil {
		log.Fatal("error occurred while auth service " + err.Error())
	}

	subscriber.Unsubscribe()
}
