package handlers

import (
	"github.com/Brandon689/echo-vite/handlers/api"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/user", api.GetJSONData)
}
