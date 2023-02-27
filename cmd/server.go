package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/inzkawka/go-ecommerce/internal/warehouse/ports"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RunServer() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-c
		cancel()
	}()

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	v1 := e.Group("/v1")
	mountV1Endpoints(v1)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)
	<-q
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func mountV1Endpoints(v1 *echo.Group) {
	productsV1 := v1.Group("/products")

	productsCtrl := ports.NewProductsController()

	productsV1.POST("", productsCtrl.CreateProduct)
	productsV1.GET("/:id", productsCtrl.GetProduct)
}
