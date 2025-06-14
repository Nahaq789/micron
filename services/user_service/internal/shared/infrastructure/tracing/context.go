package tracing

import (
	"context"
	"log/slog"
)

func GetTraceIdFromCtx(ctx context.Context, k string) (TraceID, error) {
	id := ctx.Value(k).(string)
	traceId, err := NewTraceID(id)
	if err != nil {
		slog.ErrorContext(ctx, "トレースIDの取得に失敗しました。", "error", err)
		return TraceID{}, err
	}
	return traceId, err
}
