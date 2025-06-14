package tracing

import "errors"

type TracingID interface {
	GetKey() string
	GetValue() string
}

type TraceID struct {
	key   TracingID
	value string
}

func (t TraceID) GetKey() string {
	return t.key.GetContextKey()
}

func (t TraceID) GetTraceID() string {
	return t.value
}

func NewTraceID(v string) (TraceID, error) {
	if err := validateTraceID(v); err != nil {
		return TraceID{}, err
	}

	return TraceID{value: v}, nil
}

func GenerateTraceID() TraceID {
	u := generateTraceID()
	return TraceID{value: u}
}

func validateTraceID(v string) error {
	if len(v) == 0 {
		return errors.New("トレースIDが空白です。")
	}
	return nil
}
