package ports

import (
	"net/http"

	"github.com/inzkawka/go-ecommerce/internal/warehouse/adapters"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/command"
	"github.com/labstack/echo/v4"
)

type ProductsController struct {
	app *app.Application
}

func NewProductsController() *ProductsController {
	app, err := app.NewApplication(adapters.NewInMemoryProductsRepository())
	if err != nil {
		panic(err)
	}

	return &ProductsController{
		app: app,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreatedResponse struct {
	ID string `json:"id"`
}

func (r *ProductsController) CreateProduct(e echo.Context) error {
	var cmd command.CreateProduct

	err := e.Bind(&cmd)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	id, err := r.app.Commands.CreateProduct.Handle(cmd)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, CreatedResponse{
		ID: id.String(),
	})
}
