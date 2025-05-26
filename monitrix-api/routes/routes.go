package routes

import (
	"monitrix-api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/registro", controllers.RegisterUser)
}
