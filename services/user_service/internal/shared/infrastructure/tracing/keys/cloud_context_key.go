package keys

import (
	"context"

	"github.com/google/uuid"
)

type CloudContextKey string

const (
	cloudTraceKey CloudContextKey = "X-Amzn-Trace-Id"
)

type CloudContextTrace struct {
	key CloudContextKey
}

func NewCloudContextTrace() CloudContextTrace {
	return CloudContextTrace{key: cloudTraceKey}
}

func (c CloudContextTrace) GetKey() string {
	return string(c.key)
}

func (c CloudContextTrace) GetValueFromCtx(ctx context.Context) string {
	if id, ok := ctx.Value(string(c.key)).(string); ok && id != "" {
		return id
	}
	return ""
}

func (c CloudContextTrace) GenerateID() string {
	id := uuid.New().String()
	return id
}
