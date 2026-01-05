package server

import (
	"net/http"

	appErrors "github.com/jabeedhexanovamedia/echo-starter/internal/errors"
	"github.com/jabeedhexanovamedia/echo-starter/internal/observability"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

//? USAGE:
// NewHTTPErrorHandler creates a new HTTP error handler with logging and error response formatting.
// if err := c.Validate(&req); err != nil {
// 	return errors.BadRequest("Invalid request payload", err)
// }
// OR
//if user == nil {
// 	return errors.NotFound("User not found", nil)
// }

func NewHTTPErrorHandler(log zerolog.Logger, env string) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var (
			status = http.StatusInternalServerError
			code   = "INTERNAL_ERROR"
			msg    = "Something went wrong"
		)

		// App error
		if appErr, ok := err.(*appErrors.AppError); ok {
			status = appErr.StatusCode
			code = appErr.Code
			msg = appErr.Message
		}

		// Echo HTTP error (bind / validation / not found)
		if httpErr, ok := err.(*echo.HTTPError); ok {
			status = httpErr.Code
			code = "HTTP_ERROR"
			msg = http.StatusText(httpErr.Code)
		}

		logger := observability.LoggerWithRequestID(log, c)

		// Log full error only in non-client errors
		if status >= 500 {
			logger.Error().
				Err(err).
				Int("status", status).
				Msg("server_error")
		} else {
			logger.Warn().
				Int("status", status).
				Msg("client_error")
		}

		// Hide internal details in production
		if env == "production" && status == http.StatusInternalServerError {
			msg = "Internal server error"
		}

		_ = c.JSON(status, appErrors.APIError{
			Code:    code,
			Message: msg,
		})
	}
}
