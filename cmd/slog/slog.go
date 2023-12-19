package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/art-pub/golang-slog-grafana/pkg/stats"
)

const UPDATE_IN_SECONDS = 60

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	for {
		stats.UpdateStats(logger)
		time.Sleep((UPDATE_IN_SECONDS * time.Second))
	}
}
