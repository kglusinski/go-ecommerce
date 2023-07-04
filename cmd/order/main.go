package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
	"tracing_example/internal/common/config"
	"tracing_example/internal/common/tracing"
	"tracing_example/internal/order"
	"tracing_example/internal/order/server"
)

var (
	cfg = config.InitConfig()
)

func main() {
	log.Printf("[order] run service with config: %+v", cfg)
	tracing.InitTracer(cfg)
	defer sentry.Flush(2 * time.Second)

	svc := server.StartServer(orchestrateApplication())

	svc.Logger.Fatal(svc.Start(":8001"))
}

func orchestrateApplication() *order.App {
	repo := order.NewInMemoryOrderRepository()
	return order.NewApp(repo)
}
