package main

import (
	"log/slog"
	"net/http"
	"os"
	"user_service/internal/shared/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	contextHandler := logger.NewContextHandler(handler)
	logger := slog.New(contextHandler)
	slog.SetDefault(logger)

	r := gin.Default()
	if err := Router(r); err != nil {
		slog.Error("failed to start router", "error", err)
	}
	http.ListenAndServe(":8080", r)
}
