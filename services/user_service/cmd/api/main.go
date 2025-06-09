package main

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	r := gin.Default()
	if err := Router(ctx, r); err != nil {
		slog.Error("failed to start router", "error", err)
	}
	r.Run()
}
