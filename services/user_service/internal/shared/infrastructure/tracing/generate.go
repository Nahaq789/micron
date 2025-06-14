package tracing

import "github.com/google/uuid"

func generateTraceID() string {
	u := uuid.New().String()
	return u
}
