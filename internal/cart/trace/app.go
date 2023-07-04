package trace

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"tracing_example/internal/cart"
)

type TracedApp struct {
	inner *cart.App
}

func NewTracedApp(inner *cart.App) *TracedApp {
	return &TracedApp{
		inner,
	}
}

func (a *TracedApp) CreateNewCart(ctx context.Context) uuid.UUID {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "cart.create_new_cart")
	defer span.End()

	return a.inner.CreateNewCart(newCtx)
}

func (a *TracedApp) AddItemToCart(ctx context.Context, id uuid.UUID, item cart.Item) {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "cart.add_item_to_cart")
	defer span.End()

	a.inner.AddItemToCart(newCtx, id, item)
}

func (a *TracedApp) GetCartItems(ctx context.Context, id uuid.UUID) *cart.Cart {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "cart.get_cart_items")
	defer span.End()

	return a.inner.GetCartItems(newCtx, id)
}

func (a *TracedApp) PlaceOrder(ctx context.Context, uid uuid.UUID) (uuid.UUID, error) {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "cart.place_order")
	defer span.End()

	return a.inner.PlaceOrder(newCtx, uid)
}
