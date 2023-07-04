package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	trace2 "go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"
	"tracing_example/internal/cart"
	"tracing_example/internal/cart/trace"
	"tracing_example/internal/common/config"
	"tracing_example/internal/common/tracing"
)

var (
	tracedApp *trace.TracedApp
	tracer    = otel.Tracer("cart_service")
	cfg       = config.InitConfig()
)

func main() {
	log.Printf("[Cart] run service with config: %+v", cfg)
	tracing.InitTracer(cfg)
	defer sentry.Flush(2 * time.Second)

	e := echo.New()

	e.Use(otelecho.Middleware("Cart Service"))

	orderClient := cart.NewOrderClient()
	tracedOrderClient := trace.NewTracedOrderClient(orderClient)
	repo := cart.NewInMemoryCartRepository()
	tracedRepo := trace.NewTracedCartRepository(repo)
	app := cart.NewApp(tracedOrderClient, tracedRepo)
	tracedApp = trace.NewTracedApp(app)

	e.POST("/v1/cart", CreateCart)
	e.GET("/v1/cart/:id", GetCartItems)
	e.PUT("/v1/cart/:id", AddThingToCart)
	e.POST("/v1/cart/:id", PlaceOrder)

	e.Logger.Fatal(e.Start(":8000"))
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

func CreateCart(e echo.Context) error {
	ctx := e.Request().Context()

	newCtx, span := tracer.Start(ctx, "POST /v1/cart")
	defer span.End()

	log.Println("creating cart...")
	id := tracedApp.CreateNewCart(newCtx)

	return e.JSON(http.StatusCreated, CreateResponse{ID: id})
}

func GetCartItems(e echo.Context) error {
	ctx := e.Request().Context()

	newCtx, span := tracer.Start(ctx, "GET /v1/cart/{id}",
		trace2.WithAttributes(
			attribute.String("id", e.Param("id")),
		),
	)
	defer span.End()

	log.Printf("get cart id %s\n", e.Param("id"))

	uid, _ := uuid.Parse(e.Param("id"))
	res := tracedApp.GetCartItems(newCtx, uid)

	return e.JSON(http.StatusOK, res)
}

func AddThingToCart(e echo.Context) error {
	ctx := e.Request().Context()

	newCtx, span := tracer.Start(ctx, "PUT /v1/cart/{id}",
		trace2.WithAttributes(
			attribute.String("id", e.Param("id")),
		),
	)
	defer span.End()

	log.Printf("adding item to cart id %s\n", e.Param("id"))

	var item cart.Item
	err := e.Bind(&item)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	uid, _ := uuid.Parse(e.Param("id"))
	tracedApp.AddItemToCart(newCtx, uid, item)

	return e.JSON(http.StatusNoContent, nil)
}

func PlaceOrder(e echo.Context) error {
	ctx := e.Request().Context()

	newCtx, span := tracer.Start(ctx, "POST /v1/cart/{id}",
		trace2.WithAttributes(
			attribute.String("id", e.Param("id")),
		),
	)
	defer span.End()

	log.Printf("placing order from cart id %s\n", e.Param("id"))

	uid, _ := uuid.Parse(e.Param("id"))
	orderId, err := tracedApp.PlaceOrder(newCtx, uid)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, CreateResponse{ID: orderId})
}
