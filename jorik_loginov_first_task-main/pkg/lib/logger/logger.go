package logger

import (
	"github.com/akayo16/jorik_loginov_first_task/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
)

const (
	localDevelopmentEnvironment = "local"
	devDevelopmentEnvironment   = "dev"
	prodDevelopmentEnvironment  = "prod"
)

type Logger struct {
	log *slog.Logger
}

func Init(developmentEnvironment string) {
	logger := setupLogger(developmentEnvironment)

	slog.SetDefault(logger)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case localDevelopmentEnvironment:
		log = setupPrettySlog()
	case devDevelopmentEnvironment:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			}),
		)
	case prodDevelopmentEnvironment:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
