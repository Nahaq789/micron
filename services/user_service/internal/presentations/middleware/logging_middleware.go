package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type requestInfo struct {
	status                                            int
	contents_length                                   int64
	method, path, sourceIP, query, user_agent, errors string
	elapsed                                           time.Duration
}

func LoggingMiddleware() gin.HandlerFunc {
	start := time.Now()
	return func(c *gin.Context) {
		r := &requestInfo{
			status:          c.Writer.Status(),
			contents_length: c.Request.ContentLength,
			path:            c.Request.URL.Path,
			sourceIP:        c.ClientIP(),
			query:           c.Request.URL.RawQuery,
			user_agent:      c.Request.UserAgent(),
			errors:          c.Errors.ByType(gin.ErrorTypePrivate).String(),
			elapsed:         time.Since(start),
		}
		slog.InfoContext(c.Request.Context(), "Request Info", "Request", r.logValue())
	}
}

func (r *requestInfo) logValue() slog.Value {
	return slog.GroupValue(
		slog.Int("status", r.status),
		slog.Int64("Content-length", r.contents_length),
		slog.String("method", r.method),
		slog.String("path", r.path),
		slog.String("sourceIP", r.sourceIP),
		slog.String("query", r.query),
		slog.String("user_agent", r.user_agent),
		slog.String("errors", r.errors),
		slog.Duration("elapsed", r.elapsed),
	)
}
