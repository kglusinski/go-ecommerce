package ports

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/adapters"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/command"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/app/query"
	"github.com/inzkawka/go-ecommerce/internal/warehouse/domain"
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

type GetProductResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

func (r *ProductsController) GetProduct(e echo.Context) error {
	id, err := uuid.Parse(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	product, err := r.app.Queries.GetProduct.Handle(query.GetSingleProduct{ID: id})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, toGetProductResponse(product))
}

func toGetProductResponse(product *domain.Product) *GetProductResponse {
	return &GetProductResponse{
		ID:     product.ID().String(),
		Name:   product.Name(),
		Price:  product.Price(),
		Amount: product.Amount(),
	}
}
