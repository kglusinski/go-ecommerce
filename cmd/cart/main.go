package main

import (
	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"log"
	"net/http"
	"time"
	"tracing_example/internal/cart"
	"tracing_example/internal/cart/server"
	"tracing_example/internal/cart/trace"
	"tracing_example/internal/common/config"
	"tracing_example/internal/common/tracing"
)

var (
	cfg = config.InitConfig()
)

func main() {
	log.Printf("[Cart] run service with config: %+v", cfg)
	tracing.InitTracer(cfg)
	defer sentry.Flush(2 * time.Second)

	svc := server.StartServer(orchestrateApplication())

	svc.Logger.Fatal(svc.Start(":8000"))
}

func orchestrateApplication() cart.AppInterface {
	httpClient := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	orderClient := cart.NewOrderClient(httpClient)
	tracedOrderClient := trace.NewTracedOrderClient(orderClient)
	repo := cart.NewInMemoryCartRepository()
	tracedRepo := trace.NewTracedCartRepository(repo)
	app := cart.NewApp(tracedOrderClient, tracedRepo)
	return trace.NewTracedApp(app)
}
