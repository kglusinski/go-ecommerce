package cart

import (
	"context"
	"github.com/google/uuid"
)

type AppInterface interface {
	CreateNewCart(ctx context.Context) uuid.UUID
	AddItemToCart(ctx context.Context, id uuid.UUID, item Item)
	GetCartItems(ctx context.Context, id uuid.UUID) *Cart
	PlaceOrder(ctx context.Context, uid uuid.UUID) (uuid.UUID, error)
}

type App struct {
	os   OrderService
	repo CartRepository
}

func NewApp(os OrderService, repo CartRepository) *App {
	return &App{
		os:   os,
		repo: repo,
	}
}

func (a *App) CreateNewCart(ctx context.Context) uuid.UUID {
	cart := NewCart(a.os)

	a.repo.New(ctx, cart)

	return cart.ID
}

func (a *App) AddItemToCart(ctx context.Context, id uuid.UUID, item Item) {
	cart := a.repo.Get(ctx, id)

	cart.AddItem(item)

	a.repo.Update(ctx, cart)
}

func (a *App) GetCartItems(ctx context.Context, id uuid.UUID) *Cart {
	return a.repo.Get(ctx, id)
}

func (a *App) PlaceOrder(ctx context.Context, uid uuid.UUID) (uuid.UUID, error) {
	cart := a.repo.Get(ctx, uid)

	oid, err := cart.Finish(ctx)

	a.repo.Delete(ctx, uid)
	return oid, err
}
