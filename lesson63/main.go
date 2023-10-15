package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("This is log")
	slog.Debug("This is Debug Level")
	slog.Info("This is Info Level")
	slog.Warn("This is Warn Level")
	slog.Error("This is Error Level")

	slog.Log(context.Background(), -4, "This is Debug Level")
	slog.Log(context.Background(), 0, "This is Info Level")
	slog.Log(context.Background(), 4, "This is Warn Level")
	slog.Log(context.Background(), 8, "This is Error Level")

	slog.Info("This is Info Level", "uid", 1001)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("This is a new text Info Level", "uid", 1002)

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("This is a json format Info Level", "uid", 1003)

	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	textLogger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	textLogger.Debug("This is a logger of support Debug Level", "uid", 1004)
}
