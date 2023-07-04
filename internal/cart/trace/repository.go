package trace

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"tracing_example/internal/cart"
)

type TracedCartRepository struct {
	inner cart.CartRepository
}

func NewTracedCartRepository(inner cart.CartRepository) *TracedCartRepository {
	return &TracedCartRepository{
		inner,
	}
}

func (r *TracedCartRepository) New(ctx context.Context, cart *cart.Cart) {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "repository.new")
	defer span.End()

	r.inner.New(newCtx, cart)
}

func (r *TracedCartRepository) Get(ctx context.Context, id uuid.UUID) *cart.Cart {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "repository.get")
	defer span.End()

	return r.inner.Get(newCtx, id)
}

func (r *TracedCartRepository) Update(ctx context.Context, cart *cart.Cart) {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "repository.delete")
	defer span.End()

	r.inner.Update(newCtx, cart)
}

func (r *TracedCartRepository) Delete(ctx context.Context, id uuid.UUID) {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "repository.delete")
	defer span.End()

	r.inner.Delete(newCtx, id)
}
