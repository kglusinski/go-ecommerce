package trace

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"tracing_example/internal/cart"
)

type TracedHttpOrderClient struct {
	inner *cart.HttpOrderClient
}

func NewTracedOrderClient(inner *cart.HttpOrderClient) *TracedHttpOrderClient {
	return &TracedHttpOrderClient{
		inner,
	}
}

func (c *TracedHttpOrderClient) MakeOrder(ctx context.Context, items []cart.Item) (uuid.UUID, error) {
	newCtx, span := otel.Tracer("cart_service").Start(ctx, "order_client.make_order")
	defer span.End()

	return c.inner.MakeOrder(newCtx, items)
}
