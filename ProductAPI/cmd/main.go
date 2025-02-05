package main

import (
	"product-api/Database"
	"product-api/Routs"

	"github.com/labstack/echo/v4"
)

func main() {

	Database.ConnectDatabase()

	Database.MigrateDatabase()

	e := echo.New()

	Routs.RegisterRoute(e)

	e.Logger.Fatal(e.Start(":8081"))

}