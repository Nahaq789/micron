package keys

type CloudContextKey string

const (
	cloudTraceKey CloudContextKey = "X-Amzn-Trace-Id"
)

type CloudContextTraceID struct {
	key   CloudContextKey
	value string
}

func (c CloudContextTraceID) GetKey() string {
	return string(c.key)
}

func (c CloudContextTraceID) GetValue() string {
	return c.value
}

func NewCloudContextTraceID(v string) CloudContextTraceID {
	return CloudContextTraceID{
		key:   cloudTraceKey,
		value: v,
	}
}
