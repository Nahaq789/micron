package userprofile

import (
	"testing"
)

func TestNewBio(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "normal bio",
			input: "I am a software engineer",
		},
		{
			name:  "empty bio",
			input: "",
		},
		{
			name:  "long bio",
			input: "This is a very long bio that contains a lot of information about the user including their interests, hobbies, and professional background",
		},
		{
			name:  "bio with special characters",
			input: "Hello! 🌟 I'm a developer & designer 💻 #coding",
		},
		{
			name:  "bio with newlines",
			input: "Line 1\nLine 2\nLine 3",
		},
		{
			name:  "bio with numbers",
			input: "Born in 1990, working since 2015",
		},
		{
			name:  "single character bio",
			input: "X",
		},
		{
			name:  "bio with whitespace",
			input: "  Spaces around  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewBio(tt.input)

			if result.value != tt.input {
				t.Errorf("NewBio() result = %q, want %q", result.value, tt.input)
			}
		})
	}
}
