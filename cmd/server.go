package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/inzkawka/go-ecommerce/internal/warehouse/app"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/ports"
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

	router := ports.NewRouter()

	if err := http.ListenAndServe(":8080", router.Echo); err != nil {
		log.Fatal("shutting down the server...")
	}
}

func setupApplication() (*app.Application, error) {
	return app.NewApplication(nil)
}
