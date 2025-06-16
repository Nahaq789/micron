package tracing

type TracingID interface {
	GetKey() string
	GetValueFromCtx() string
	GenerateID() string
}
