package main

import (
	"log/slog"
	"user_service/internal/presentations/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r gin.IRouter) error {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())
	v1.GET("/ping", func(c *gin.Context) {
		slog.InfoContext(c.Request.Context(), "良いこのみんな～")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return nil
}
