package main

import (
	"github.com/gilxa1226/microservices/order/config"
	"github.com/gilxa1226/microservices/order/internal/adapters/db"
	"github.com/gilxa1226/microservices/order/internal/adapters/grpc"
	"github.com/gilxa1226/microservices/order/internal/application/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
