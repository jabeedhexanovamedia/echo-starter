package server

import (
	"fmt"

	"github.com/jabeedhexanovamedia/echo-starter/internal/config"
	"github.com/jabeedhexanovamedia/echo-starter/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(cfg *config.Config) *echo.Echo {
	e := echo.New()

	// Global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	e.GET("/", handler.Health)

	fmt.Printf("Server starting in %s mode on port %s\n", cfg.Server.Env, cfg.Server.Port)
	return e
}
