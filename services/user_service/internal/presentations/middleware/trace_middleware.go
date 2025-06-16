package middleware

import (
	"context"
	"user_service/internal/shared/infrastructure/tracing/keys"

	"github.com/gin-gonic/gin"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cloudTrace := keys.NewCloudContextTrace()
		contextTrace := keys.NewContextTrace()
		cloudTraceID := c.Request.Header.Get(cloudTrace.GetKey())
		if cloudTraceID == "" {
			cloudTraceID = cloudTrace.GenerateID()
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, cloudTrace.GetKey(), cloudTraceID)
		ctx = context.WithValue(ctx, contextTrace.GetKey(), contextTrace.GenerateID())

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
