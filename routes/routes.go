package routes

import (
	"auth_service/controllers"

	"github.com/labstack/echo/v4"
)

func RegistRoutes(e *echo.Echo) {
	e.POST("/api/login", controllers.Login)
}
