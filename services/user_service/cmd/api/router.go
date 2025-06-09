package main

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Router(ctx context.Context, r gin.IRouter) error {
	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return nil
}
