package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yandzee/wait-action/internal/app"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if err := app.Run(context.Background(), logger); err != nil {
		logger.Error("application run failed", "err", err.Error())
		os.Exit(1)
	}
}
