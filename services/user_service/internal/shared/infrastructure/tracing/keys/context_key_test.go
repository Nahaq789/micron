package keys

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestNewContextTrace(t *testing.T) {
	trace := NewContextTrace()

	// Verify the key is set correctly
	if trace.key != contextTraceKey {
		t.Errorf("Expected key to be %q, got %q", contextTraceKey, trace.key)
	}

	// Verify the key value is correct
	expectedKey := "traceID"
	if string(trace.key) != expectedKey {
		t.Errorf("Expected key value to be %q, got %q", expectedKey, string(trace.key))
	}
}

func TestContextTrace_GetKey(t *testing.T) {
	trace := NewContextTrace()
	key := trace.GetKey()

	expected := "traceID"
	if key != expected {
		t.Errorf("Expected GetKey() to return %q, got %q", expected, key)
	}
}

func TestContextTrace_GetValueFromCtx(t *testing.T) {
	trace := NewContextTrace()

	tests := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			name:     "context with valid string value",
			ctx:      context.WithValue(context.Background(), string(contextTraceKey), "trace-123"),
			expected: "trace-123",
		},
		{
			name:     "context with empty string value",
			ctx:      context.WithValue(context.Background(), string(contextTraceKey), ""),
			expected: "",
		},
		{
			name:     "context without the key",
			ctx:      context.Background(),
			expected: "",
		},
		{
			name:     "context with wrong value type",
			ctx:      context.WithValue(context.Background(), string(contextTraceKey), 123),
			expected: "",
		},
		{
			name:     "context with nil value",
			ctx:      context.WithValue(context.Background(), string(contextTraceKey), nil),
			expected: "",
		},
		{
			name:     "context with different key",
			ctx:      context.WithValue(context.Background(), "different-key", "some-value"),
			expected: "",
		},
		{
			name:     "context with ContextKey type as key",
			ctx:      context.WithValue(context.Background(), contextTraceKey, "trace-456"),
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

func TestContextTrace_GenerateID(t *testing.T) {
	trace := NewContextTrace()

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

// Test helper function that could be useful for integration tests
func TestContextTrace_Integration(t *testing.T) {
	trace := NewContextTrace()

	t.Run("complete workflow", func(t *testing.T) {
		// Generate an ID
		id := trace.GenerateID()

		// Create context with the ID
		ctx := context.WithValue(context.Background(), trace.GetKey(), id)

		// Retrieve the ID from context
		retrievedID := trace.GetValueFromCtx(ctx)

		if retrievedID != id {
			t.Errorf("Expected retrieved ID to be %q, got %q", id, retrievedID)
		}
	})
}

// Benchmark tests
func BenchmarkNewContextTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewContextTrace()
	}
}

func BenchmarkContextTrace_GetKey(b *testing.B) {
	trace := NewContextTrace()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trace.GetKey()
	}
}

func BenchmarkContextTrace_GetValueFromCtx(b *testing.B) {
	trace := NewContextTrace()
	ctx := context.WithValue(context.Background(), string(contextTraceKey), "trace-123")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trace.GetValueFromCtx(ctx)
	}
}

func BenchmarkContextTrace_GenerateID(b *testing.B) {
	trace := NewContextTrace()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trace.GenerateID()
	}
}

// Example test showing how the package might be used
func ExampleContextTrace() {
	trace := NewContextTrace()

	// Generate a new trace ID
	id := trace.GenerateID()

	// Add it to context
	ctx := context.WithValue(context.Background(), trace.GetKey(), id)

	// Retrieve it later
	retrievedID := trace.GetValueFromCtx(ctx)

	if retrievedID == id {
		// Success!
	}
}
