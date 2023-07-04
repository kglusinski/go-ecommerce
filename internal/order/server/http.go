package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"tracing_example/internal/order"
)

var (
	svc    order.AppInterface
	tracer = otel.Tracer("cart_service")
)

func StartServer(app order.AppInterface) *echo.Echo {
	e := echo.New()
	svc = app

	e.Use(otelecho.Middleware("Order Service"))

	e.POST("/v1/orders", CreateOrder)

	return e
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

	id := svc.CreateNewOrder(req.Items)

	span.SetStatus(codes.Ok, "Created")
	return e.JSON(http.StatusCreated, CreateResponse{ID: id})
}
