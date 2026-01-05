package observability

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		status := c.Response().Status
		method := c.Request().Method
		path := c.Path()

		HttpRequestsTotal.WithLabelValues(
			method,
			path,
			strconv.Itoa(status),
		).Inc()

		return err
	}
}
