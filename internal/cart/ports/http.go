package ports

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/inzkawka/go-ecommerce/internal/cart/app"
	"github.com/inzkawka/go-ecommerce/internal/cart/app/command"
	"github.com/inzkawka/go-ecommerce/internal/cart/app/query"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	app *app.Application
}

func NewCartController(app *app.Application) *CartController {
	return &CartController{app: app}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreatedResponse struct {
	ID string `json:"id"`
}

func (c *CartController) CreateCart(e echo.Context) error {
	// TODO: user id should be taken from the context
	var cmd = command.CreateCart{
		UserID: uuid.New(),
	}

	id, err := c.app.Commands.CreateCart.Handle(cmd)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, CreatedResponse{
		ID: id.String(),
	})
}

func (c *CartController) AddToCart(e echo.Context) error {
	var cmd command.AddToCart

	err := e.Bind(&cmd)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	cmd.CartID, err = uuid.Parse(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	err = c.app.Commands.AddToCart.Handle(cmd)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return e.NoContent(http.StatusNoContent)
}

func (c *CartController) RemoveFromCart(e echo.Context) error {
	var cmd command.RemoveFromCart

	err := e.Bind(&cmd)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	cmd.CartID, err = uuid.Parse(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	err = c.app.Commands.RemoveFromCart.Handle(cmd)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return e.NoContent(http.StatusNoContent)
}

func (c *CartController) GetCart(e echo.Context) error {
	id, err := uuid.Parse(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	cart, err := c.app.Queries.GetCart.Handle(&query.GetCart{
		ID: id,
	})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, cart)
}
