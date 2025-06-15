package keys

type ContextKey string

const (
	ContextTraceKey ContextKey = "traceID"
)

func (c ContextKey) GetContextKey() string {
	return string(ContextTraceKey)
}

type ContextTraceID struct {
	key   ContextKey
	value string
}

func (c ContextTraceID) GetKey() string {
	return string(c.key)
}

func (c ContextTraceID) GetValue() string {
	return c.value
}

func NewContextTraceID(v string) ContextTraceID {
	return ContextTraceID{
		key:   ContextTraceKey,
		value: v,
	}
}
