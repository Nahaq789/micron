package logger

import (
	"context"
	"log/slog"
	"user_service/internal/shared/infrastructure/tracing/keys"
)

type ContextHandler struct {
	handler slog.Handler
}

func NewContextHandler(handler slog.Handler) *ContextHandler {
	return &ContextHandler{handler: handler}
}

func (h *ContextHandler) Handle(ctx context.Context, record slog.Record) error {
	var attrs []slog.Attr

	cloudTrace := keys.NewCloudContextTrace()
	trace := keys.NewContextTrace()

	if id := cloudTrace.GetValueFromCtx(ctx); id != "" {
		attrs = append(attrs, slog.String(cloudTrace.GetKey(), id))
	}
	if id := trace.GetValueFromCtx(ctx); id != "" {
		attrs = append(attrs, slog.String(trace.GetKey(), id))
	}

	if len(attrs) > 0 {
		record.AddAttrs(attrs...)
	}

	return h.handler.Handle(ctx, record)
}

func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewContextHandler(h.handler.WithAttrs(attrs))
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return NewContextHandler(h.handler.WithGroup(name))
}
