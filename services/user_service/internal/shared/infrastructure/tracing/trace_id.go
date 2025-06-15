package tracing

type TracingID interface {
	GetKey() string
	GetValue() string
}
