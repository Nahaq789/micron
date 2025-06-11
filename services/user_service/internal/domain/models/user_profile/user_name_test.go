package userprofile

import (
	"testing"
)

func TestNewUserName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid username",
			input:   "testuser",
			wantErr: false,
		},
		{
			name:    "single character username",
			input:   "a",
			wantErr: false,
		},
		{
			name:    "long username",
			input:   "verylongusernamewithalotofcharacters",
			wantErr: false,
		},
		{
			name:    "username with numbers",
			input:   "user123",
			wantErr: false,
		},
		{
			name:    "username with special characters",
			input:   "user_name-123",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewUserName(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewUserName() error = nil, wantErr %v", tt.wantErr)
				}
				if result.value != "" {
					t.Errorf("NewUserName() result should be empty on error")
				}
			} else {
				if err != nil {
					t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				}
				if result.value != tt.input {
					t.Errorf("NewUserName() result = %v, want %v", result.value, tt.input)
				}
			}
		})
	}
}

func TestNewUserName_ErrorMessage(t *testing.T) {
	_, err := NewUserName("")
	if err == nil {
		t.Error("Expected error for empty username")
	}
	expected := "ユーザ名は1文字以上にしてください。"
	if err.Error() != expected {
		t.Errorf("Expected error message %q, got %q", expected, err.Error())
	}
}

func TestValidateUserName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid username",
			input:   "validuser",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   "x",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUserName(tt.input)

			if tt.wantErr && err == nil {
				t.Errorf("validateUserName() error = nil, wantErr %v", tt.wantErr)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("validateUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
