package main

import (
	"context"
	"log/slog"
	"os"
	"user_service/internal/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	contextHandler := logger.NewContextHandler(handler)
	logger := slog.New(contextHandler)
	slog.SetDefault(logger)

	r := gin.Default()
	if err := Router(ctx, r); err != nil {
		slog.Error("failed to start router", "error", err)
	}
	r.Run()
}
