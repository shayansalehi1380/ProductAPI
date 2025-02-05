package Handlers

import (
	"net/http"
	"product-api/Database"
	"product-api/Models"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {

	var products []Models.Product

	Database.DB.Find(&products)
	return c.JSON(http.StatusOK, products)
}


func CreateProduct(c echo.Context) error {

	product := new(Models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "Create Faild")
	}
	Database.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}


func UpdateProduct(c echo.Context) error {

	id := c.Param("id")
	var product Models.Product

	if result := Database.DB.First(&product, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, "Product Not Found!")
	}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, "Update Faild")
	}
	Database.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}


func DeleteProduct(c echo.Context) error {
    id := c.Param("id")
    var product Models.Product

    // پیدا کردن محصول بر اساس ID
    if result := Database.DB.First(&product, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, "Product Not Found!")
    }

    // حذف محصول
    if result := Database.DB.Delete(&product); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, "Delete Failed")
    }

    return c.JSON(http.StatusOK, product)
}