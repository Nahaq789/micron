package tracing

import (
	"context"
)

func GetTraceIdFromCtx(ctx context.Context, c TracingID) string {
	id := ctx.Value(c.GetKey()).(string)
	return id
}
