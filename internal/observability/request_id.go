package observability

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

const RequestIDKey = "request_id"

func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)
			if reqID == "" {
				reqID = c.Request().Header.Get(echo.HeaderXRequestID)
			}

			if reqID == "" {
				reqID = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			c.Set(RequestIDKey, reqID)
			return next(c)
		}
	}
}

func LoggerWithRequestID(log zerolog.Logger, c echo.Context) zerolog.Logger {
	reqID, ok := c.Get(RequestIDKey).(string)
	if ok && reqID != "" {
		return log.With().Str("request_id", reqID).Logger()
	}
	return log
}
