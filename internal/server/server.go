package server

import (
	"github.com/jabeedhexanovamedia/echo-starter/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	e.GET("/", handler.Health)

	return e
}
