package stats

import (
	"log/slog"
	"math/rand"
)

func UpdateStats(logger *slog.Logger) {
	logger.Info("Usage Statistics",
		slog.Int("current-memory", rand.Intn(60)+20),
		slog.Int("min-memory", 20),
		slog.Int("max-memory", 80),
		slog.Int("cpu", rand.Intn(16)+1),
	)
}
