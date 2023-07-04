package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"tracing_example/internal/cart"
)

var (
	svc    cart.AppInterface
	tracer = otel.Tracer("cart_service")
)

func StartServer(app cart.AppInterface) *echo.Echo {
	e := echo.New()
	svc = app

	e.Use(otelecho.Middleware("Cart Service"))

	e.POST("/v1/cart", CreateCart)
	e.GET("/v1/cart/:id", GetCartItems)
	e.PUT("/v1/cart/:id", AddThingToCart)
	e.POST("/v1/cart/:id", PlaceOrder)

	return e
}

type CreatedResponse struct {
	ID uuid.UUID `json:"id"`
}
