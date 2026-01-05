package observability

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func NewLogger(env string, level string) zerolog.Logger {
	logLevel := zerolog.InfoLevel

	if parsed, err := zerolog.ParseLevel(strings.ToLower(level)); err == nil {
		logLevel = parsed
	}

	zerolog.SetGlobalLevel(logLevel)

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("env", env).
		Logger()

	// Pretty logs only in development
	if env == "development" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return logger
}

// ? USAGE
// This project uses STRUCTURED LOGGING (zerolog).

//  BASIC LOG LEVELS

// logger.Debug().Msg("debug message") // development diagnostics
// logger.Info().Msg("info message") // normal app flow
// logger.Warn().Msg("warning message") // recoverable issues
// logger.Error().Msg("error message") // failures
// logger.Fatal().Msg("fatal message") // app crash (exits)

// LOG WITH STRUCTURED FIELDS (RECOMMENDED)
// BAD:
// logger.Info().Msg("server started on port " + port)

// GOOD:
// logger.Info().
// Str("port", port).
// Str("env", env).
// Msg("server started")

// LOGGING ERRORS (ALWAYS USE Err)

// BAD:
// logger.Error().Msg(err.Error())

// GOOD:
// logger.Error().
// Err(err).
// Msg("failed to connect to database")

// HANDLER LOGGING (WITH REQUEST ID)

// Inside HTTP handlers, always use request-aware logger:

// log := observability.LoggerWithRequestID(logger, c)

// log.Info().
// Str("user_id", userID).
// Msg("user created")

// This automatically adds:

// - request_id
// - method
// - path

// SERVICE / REPOSITORY LOGGING

// Use normal logger (no HTTP context):

// logger.Info().
// Str("email", email).
// Msg("creating user in database")
