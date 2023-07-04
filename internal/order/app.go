package order

import "github.com/google/uuid"

type AppInterface interface {
	CreateNewOrder(items []Product) uuid.UUID
}

type OrderRepository interface {
	New(order *Order)
}

type App struct {
	repo OrderRepository
}

func NewApp(repo OrderRepository) *App {
	return &App{repo: repo}
}

func (a *App) CreateNewOrder(items []Product) uuid.UUID {
	order := NewOrder(items)

	a.repo.New(order)

	return order.ID
}
