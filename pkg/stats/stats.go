package stats

import "log/slog"

func UpdateStats(logger *slog.Logger) {
	logger.Info("Usage Statistics",
		slog.Int("current-memory", 50),
		slog.Int("min-memory", 20),
		slog.Int("max-memory", 80),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
}
