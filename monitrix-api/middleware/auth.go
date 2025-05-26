package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("auth_token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Acesso não autorizado",
			})
		}
		return next(c)
	}
}
