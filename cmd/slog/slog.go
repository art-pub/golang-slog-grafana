package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/art-pub/golang-slog-grafana/pkg/stats"
)

// how often do we want to update the log?
const UPDATE_IN_SECONDS = 5

// where should the log be saved?
const LOGFILE = "../../logs/slog.log"

func main() {

	f, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err.Error())
	}
	defer f.Close()

	logger := slog.New(slog.NewJSONHandler(f, nil))

	for {
		stats.UpdateStats(logger)
		time.Sleep((UPDATE_IN_SECONDS * time.Second))
	}
}
