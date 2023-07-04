package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"
	"tracing_example/internal/common/config"
	"tracing_example/internal/common/tracing"
	"tracing_example/internal/order"
)

var (
	app    *order.App
	tracer = otel.Tracer("order_service")
	cfg    = config.InitConfig()
)

func main() {
	tracing.InitTracer(cfg)

	defer sentry.Flush(2 * time.Second)
	e := echo.New()

	repo := order.NewInMemoryOrderRepository()
	app = order.NewApp(repo)

	e.Use(otelecho.Middleware("Order Service"))

	e.POST("/v1/orders", CreateOrder)

	e.Logger.Fatal(e.Start(":8001"))
}

type CreateRequest struct {
	Items []order.Product `json:"items"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

func CreateOrder(e echo.Context) error {
	ctx := e.Request().Context()
	span := trace.SpanFromContext(ctx)
	if span == nil {
		_, span = tracer.Start(ctx, "POST /v1/orders")
	}
	defer span.End()

	log.Println("placing order...")
	log.Println(span)

	var req CreateRequest
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	id := app.CreateNewOrder(req.Items)

	span.SetStatus(codes.Ok, "Created")
	return e.JSON(http.StatusCreated, CreateResponse{ID: id})
}
