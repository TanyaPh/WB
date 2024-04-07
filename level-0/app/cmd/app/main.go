package main

import (
	"api/internal/cache"
	"api/internal/controller"
	"api/internal/repository"
	"api/internal/service"
	"api/pkg/postgres"
	"api/pkg/server"
	"context"
	"fmt"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	// "github.com/spf13/viper"
)

func main() {
	pg, err := postgres.New(context.Background(), 
							"host=localhost dbname=demo_service user=gaby password=forza sslmode=disable")
	if err != nil {
		logrus.Fatalf("Unable to connection to database: %v\n", err)
	}
	logrus.Info("Successfully connected to DB")

	repos := repository.NewRepository(pg)
	cache, err := cache.NewCache(repos)
	if err != nil {
		logrus.Fatalf("Unable to create cache: %v\n", err)
	}
	logrus.Info("Successfully create cache")

	services := service.NewService(repos, cache)
	handlers := controller.NewRouter(services)
	srv := server.New("8080", handlers)

	sc, err := stan.Connect("demo-cluster", "demo-subscribe", stan.NatsURL("nats://localhost:1234"))
	if err != nil {
		logrus.Fatalf("Unable to connection to STAN: %v\n", err)
	}
	logrus.Info("Successfully connected to the STAN")
	defer sc.Close()

	sc.Subscribe("OrderChannel", func(msg *stan.Msg) {
		fmt.Println(msg.Data)
		// services.Order.Create(msg)
	}, stan.DurableName("OrderChannelForever"))
	logrus.Info("Successfully Subscribe to the OrderChannel")

	logrus.Info("the server starting")
	if err := srv.Run(); err != nil {
		logrus.Fatal(err)
	}
	
}
