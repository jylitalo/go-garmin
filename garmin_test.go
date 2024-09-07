package garmin

import (
	"log/slog"
)

func init() {
	loadEnv(".env")
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetLogLoggerLevel(slog.LevelInfo)
}
