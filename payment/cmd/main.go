package cmd

import (
	"github.com/gilxa1226/microservices/order/config"
	"github.com/gilxa1226/microservices/payment/internal/adapters/db"
	"github.com/gilxa1226/microservices/payment/internal/adapters/grpc"
	"github.com/gilxa1226/microservices/payment/internal/adapters/payment"
	"github.com/gilxa1226/microservices/payment/internal/application/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
