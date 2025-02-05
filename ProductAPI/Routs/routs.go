package Routs

import (
	"product-api/Handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo){

	e.GET("/products", Handlers.GetProduct)
	e.POST("/products", Handlers.CreateProduct)
	e.PUT("/products/:id", Handlers.UpdateProduct)
	e.DELETE("/products/:id", Handlers.DeleteProduct)

}