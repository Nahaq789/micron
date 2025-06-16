package keys

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string

const (
	contextTraceKey ContextKey = "traceID"
)

type ContextTrace struct {
	key ContextKey
}

func NewContextTrace() ContextTrace {
	return ContextTrace{key: contextTraceKey}
}

func (c ContextTrace) GetKey() string {
	return string(c.key)
}

func (c ContextTrace) GetValueFromCtx(ctx context.Context) string {
	if id, ok := ctx.Value(c.key).(string); ok && id != "" {
		return id
	}
	return ""
}

func (c ContextTrace) GenerateID() string {
	id := uuid.New().String()
	return id
}
