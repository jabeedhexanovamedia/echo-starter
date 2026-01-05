package server

import (
	"fmt"

	"github.com/jabeedhexanovamedia/echo-starter/internal/config"
	"github.com/jabeedhexanovamedia/echo-starter/internal/handler"
	appValidator "github.com/jabeedhexanovamedia/echo-starter/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func New(cfg *config.Config) *echo.Echo {
	e := echo.New()

	// Validator
	v := appValidator.New()
	e.Validator = appValidator.NewEchoValidator(v)

	//? USAGE: in any handler
	// if err := c.Validate(&req); err != nil {
	// 	return c.JSON(400, err)
	// }

	// Global middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// Rate Limiting
	ratePerSecond := float64(cfg.RateLimiter.Requests) / cfg.RateLimiter.Duration.Seconds()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      rate.Limit(ratePerSecond),
			Burst:     cfg.RateLimiter.Requests,
			ExpiresIn: cfg.RateLimiter.Duration,
		},
	)))

	// Register routes
	e.GET("/", handler.Health)

	fmt.Printf("Server starting in %s mode on port %s\n", cfg.Server.Env, cfg.Server.Port)
	return e
}
