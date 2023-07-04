package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	trace2 "go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"tracing_example/internal/cart"
)

type V1Handler struct{}

func CreateCart(e echo.Context) error {
	ctx := e.Request().Context()

	newCtx, span := tracer.Start(ctx, "POST /v1/cart")
	defer span.End()

	log.Println("creating cart...")
	id := svc.CreateNewCart(newCtx)

	return e.JSON(http.StatusCreated, CreatedResponse{ID: id})
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
	res := svc.GetCartItems(newCtx, uid)

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
	svc.AddItemToCart(newCtx, uid, item)

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
	orderId, err := svc.PlaceOrder(newCtx, uid)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, CreatedResponse{ID: orderId})
}
