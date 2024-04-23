package logger

import (
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/lmittmann/tint"
)

var once sync.Once
var logger *slog.Logger

func configure() *slog.Logger {
	w := os.Stderr

	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	)

	return logger
}

func GetLogger() *slog.Logger {
	once.Do(func() {
		logger = configure()
	})
	return logger
}
