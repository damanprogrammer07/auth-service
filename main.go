package main

import (
	"auth_service/configs"
	"auth_service/routes"
	"auth_service/seeds"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	configs.ConnectDB()
	seeds.SeedAdmin()

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegistRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
