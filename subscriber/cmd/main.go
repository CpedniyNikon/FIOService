package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"subscriber/internal/handler"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
	"subscriber/pkg/server"
	"time"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}

func main() {
	time.Sleep(5 * time.Second)
	fmt.Println("server started")
	if err := methods.InitConfig(); err != nil {
		logrus.Fatal("error initializing configs: %s", err.Error())
	}

	nc, err := methods.InitNatsBridge()
	if err != nil {
		logrus.Fatal("error initializing nats bridge: %s", err.Error())
	}

	db, err := methods.InitDbConnection()
	if err != nil {
		logrus.Fatal("error white connecting to db: %s", err.Error())
	}

	db.AutoMigrate(&models.PersonData{})

	handlers := handler.NewHandler()

	subscriber := methods.SubscribeModelOrder(db, nc)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("app_port"), handlers.InitAuthRoutes()); err != nil {
		logrus.Fatal("error occurred while auth service " + err.Error())
	}

	subscriber.Unsubscribe()
}
