package ports

import (
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/adapters"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/command"
	"github.com/labstack/echo/v4"
)

type Router struct {
	app  *app.Application
	Echo *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()
	app, err := app.NewApplication(adapters.NewInMemoryProductsRepository())
	if err != nil {
		panic(err)
	}

	router := &Router{
		Echo: e,
		app:  app,
	}

	e.POST("/products", router.CreateProduct)

	return router
}

func (r *Router) CreateProduct(echo.Context) error {
	cmd := command.CreateProduct{}

	return r.app.Commands.CreateProduct.Handle(cmd)
}
