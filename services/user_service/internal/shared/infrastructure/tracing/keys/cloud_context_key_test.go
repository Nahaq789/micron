package keys

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestNewCloudContextTrace(t *testing.T) {
	trace := NewCloudContextTrace()

	// Verify the key is set correctly
	if trace.key != cloudTraceKey {
		t.Errorf("Expected key to be %q, got %q", cloudTraceKey, trace.key)
	}

	// Verify the key value is correct
	expectedKey := "X-Amzn-Trace-Id"
	if string(trace.key) != expectedKey {
		t.Errorf("Expected key value to be %q, got %q", expectedKey, string(trace.key))
	}
}

func TestCloudContextTrace_GetKey(t *testing.T) {
	trace := NewCloudContextTrace()
	key := trace.GetKey()

	expected := "X-Amzn-Trace-Id"
	if key != expected {
		t.Errorf("Expected GetKey() to return %q, got %q", expected, key)
	}
}

func TestCloudContextTrace_GetValueFromCtx(t *testing.T) {
	trace := NewCloudContextTrace()

	tests := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			name:     "context with valid string value",
			ctx:      context.WithValue(context.Background(), string(cloudTraceKey), "trace-123"),
			expected: "trace-123",
		},
		{
			name:     "context with empty string value",
			ctx:      context.WithValue(context.Background(), string(cloudTraceKey), ""),
			expected: "",
		},
		{
			name:     "context without the key",
			ctx:      context.Background(),
			expected: "",
		},
		{
			name:     "context with wrong value type",
			ctx:      context.WithValue(context.Background(), string(cloudTraceKey), 123),
			expected: "",
		},
		{
			name:     "context with nil value",
			ctx:      context.WithValue(context.Background(), string(cloudTraceKey), nil),
			expected: "",
		},
		{
			name:     "context with different key",
			ctx:      context.WithValue(context.Background(), "different-key", "some-value"),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trace.GetValueFromCtx(tt.ctx)
			if result != tt.expected {
				t.Errorf("Expected GetValueFromCtx() to return %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestCloudContextTrace_GenerateID(t *testing.T) {
	trace := NewCloudContextTrace()

	t.Run("generates valid UUID", func(t *testing.T) {
		id := trace.GenerateID()

		// Verify the generated ID is not empty
		if id == "" {
			t.Error("Expected GenerateID() to return non-empty string")
		}

		// Verify it's a valid UUID format
		_, err := uuid.Parse(id)
		if err != nil {
			t.Errorf("Expected GenerateID() to return valid UUID, got error: %v", err)
		}
	})

	t.Run("generates unique IDs", func(t *testing.T) {
		id1 := trace.GenerateID()
		id2 := trace.GenerateID()

		if id1 == id2 {
			t.Errorf("Expected GenerateID() to return unique IDs, got same ID: %q", id1)
		}
	})

	t.Run("generates multiple unique IDs", func(t *testing.T) {
		ids := make(map[string]bool)
		iterations := 100

		for i := 0; i < iterations; i++ {
			id := trace.GenerateID()
			if ids[id] {
				t.Errorf("Generated duplicate ID: %q", id)
				break
			}
			ids[id] = true
		}

		if len(ids) != iterations {
			t.Errorf("Expected %d unique IDs, got %d", iterations, len(ids))
		}
	})
}

// Benchmark tests
func BenchmarkNewCloudContextTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCloudContextTrace()
	}
}

func BenchmarkCloudContextTrace_GetKey(b *testing.B) {
	trace := NewCloudContextTrace()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trace.GetKey()
	}
}

func BenchmarkCloudContextTrace_GetValueFromCtx(b *testing.B) {
	trace := NewCloudContextTrace()
	ctx := context.WithValue(context.Background(), string(cloudTraceKey), "trace-123")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trace.GetValueFromCtx(ctx)
	}
}

func BenchmarkCloudContextTrace_GenerateID(b *testing.B) {
	trace := NewCloudContextTrace()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trace.GenerateID()
	}
}
