package keys

type CloudContextKey string

const (
	cloudTraceKey CloudContextKey = "X-Amzn-Trace-Id"
)

type CloudContextTrace struct {
	key   CloudContextKey
	value string
}

func NewCloudContextTrace(v string) CloudContextTrace {
	return CloudContextTrace{
		key:   cloudTraceKey,
		value: v,
	}
}
