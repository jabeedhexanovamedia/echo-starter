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
