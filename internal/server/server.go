package server

import (
	"fmt"

	"github.com/jabeedhexanovamedia/echo-starter/internal/config"
	"github.com/jabeedhexanovamedia/echo-starter/internal/handler"
	"github.com/jabeedhexanovamedia/echo-starter/internal/observability"
	av "github.com/jabeedhexanovamedia/echo-starter/internal/validator"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func New(cfg *config.Config) *echo.Echo {
	e := echo.New()

	// Logger
	logger := observability.NewLogger(cfg.Server.Env, cfg.Logging.Level)

	// Global Error handler
	e.HTTPErrorHandler = NewHTTPErrorHandler(logger, cfg.Server.Env)

	// Observability
	observability.RegisterMetrics()

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(observability.RequestIDMiddleware())
	e.Use(observability.MetricsMiddleware)
	e.Use(middleware.Recover())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log := observability.LoggerWithRequestID(logger, c)
			log.Info().
				Str("method", c.Request().Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Err(v.Error).
				Msg("http_request")
			return nil
		},
	}))

	// Validator
	v := av.New()
	e.Validator = av.NewEchoValidator(v)

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
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	fmt.Printf("Server starting in %s mode on port %s\n", cfg.Server.Env, cfg.Server.Port)
	return e
}
