package keys

type ContextKey string

const (
	ContextTraceKey ContextKey = "traceID"
)

func NewContextKey() ContextKey {
	return ContextKey(ContextTraceKey)
}

func (c ContextKey) GetContextKey() string {
	return string(ContextTraceKey)
}
