package initialisations

import (
	"log/slog"
	"os"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/constants"
)

func InitLogger(env string) *slog.Logger {
	var logLevel slog.Level
	switch env {
	case constants.EnvDevelopment:
		logLevel = slog.LevelDebug
	case constants.EnvStaging:
		logLevel = slog.LevelInfo
	case constants.EnvProduction:
		logLevel = slog.LevelWarn
	}
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
	}))
	slog.SetDefault(jsonLogger)
	return jsonLogger
}
