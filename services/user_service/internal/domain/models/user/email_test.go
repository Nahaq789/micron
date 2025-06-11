package user

import (
	"testing"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errorMsg  string
	}{
		{
			name:      "有効なメールアドレス - 基本形",
			input:     "test@example.com",
			wantError: false,
		},
		{
			name:      "有効なメールアドレス - サブドメインあり",
			input:     "user@mail.example.com",
			wantError: false,
		},
		{
			name:      "有効なメールアドレス - 特殊文字含む",
			input:     "user.name+tag@example.com",
			wantError: false,
		},
		{
			name:      "有効なメールアドレス - 数字含む",
			input:     "user123@example123.com",
			wantError: false,
		},
		{
			name:      "有効なメールアドレス - ハイフン含む",
			input:     "test@sub-domain.example.com",
			wantError: false,
		},
		{
			name:      "無効なメールアドレス - @マークなし",
			input:     "invalidemail",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
		{
			name:      "無効なメールアドレス - ドメイン部なし",
			input:     "test@",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
		{
			name:      "無効なメールアドレス - ローカル部なし",
			input:     "@example.com",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
		{
			name:      "空文字",
			input:     "",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
		{
			name:      "無効なメールアドレス - スペース含む",
			input:     "test @example.com",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
		{
			name:      "無効なメールアドレス - 連続ドット",
			input:     "test..test@example.com",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
		{
			name:      "無効なメールアドレス - 複数@マーク",
			input:     "test@@example.com",
			wantError: true,
			errorMsg:  "メールアドレスの形式が正しくないです。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := NewEmail(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("NewEmail(%q) expected error, but got nil", tt.input)
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("NewEmail(%q) error message = %q, want %q", tt.input, err.Error(), tt.errorMsg)
				}
				if email != nil {
					t.Errorf("NewEmail(%q) expected email to be nil when error occurs, but got %v", tt.input, email)
				}
			} else {
				if err != nil {
					t.Errorf("NewEmail(%q) unexpected error: %v", tt.input, err)
					return
				}
				if email == nil {
					t.Errorf("NewEmail(%q) expected email to be non-nil, but got nil", tt.input)
					return
				}
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{
			name:      "標準的なメールアドレス",
			input:     "test@example.com",
			wantError: false,
		},
		{
			name:      "Gmail形式",
			input:     "user@gmail.com",
			wantError: false,
		},
		{
			name:      "サブドメイン付き",
			input:     "admin@mail.company.co.jp",
			wantError: false,
		},
		{
			name:      "数字を含むアドレス",
			input:     "user123@example123.com",
			wantError: false,
		},
		{
			name:      "ハイフンを含むドメイン",
			input:     "test@sub-domain.example.com",
			wantError: false,
		},
		{
			name:      "プラス記号を含む",
			input:     "user+label@example.com",
			wantError: false,
		},
		{
			name:      "アンダースコアを含む",
			input:     "user_name@example.com",
			wantError: false,
		},
		{
			name:      "ドットを含む",
			input:     "first.last@example.com",
			wantError: false,
		},
		{
			name:      "感嘆符を含む",
			input:     "user!test@example.com",
			wantError: false,
		},
		{
			name:      "パーセント記号を含む",
			input:     "user%test@example.com",
			wantError: false,
		},
		{
			name:      "@マークなし",
			input:     "testexample.com",
			wantError: true,
		},
		{
			name:      "複数の@マーク",
			input:     "test@@example.com",
			wantError: true,
		},
		{
			name:      "空文字",
			input:     "",
			wantError: true,
		},
		{
			name:      "スペースを含む",
			input:     "test @example.com",
			wantError: true,
		},
		{
			name:      "日本語を含む",
			input:     "テスト@example.com",
			wantError: true,
		},
		{
			name:      "連続するドット",
			input:     "test..name@example.com",
			wantError: true,
		},
		{
			name:      "ドメイン部が空",
			input:     "test@",
			wantError: true,
		},
		{
			name:      "ローカル部が空",
			input:     "@example.com",
			wantError: true,
		},
		{
			name:      "ドットで開始",
			input:     ".test@example.com",
			wantError: true,
		},
		{
			name:      "ドットで終了",
			input:     "test.@example.com",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateEmail(tt.input)

			if tt.wantError {
				if err == nil {
					t.Errorf("validateEmail(%q) expected error, but got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("validateEmail(%q) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

func TestEmailValidationErrorMessage(t *testing.T) {
	invalidEmails := []string{
		"testexample.com",
		"test@",
		"@example.com",
		"",
		"test..test@example.com",
	}

	expectedMsg := "メールアドレスの形式が正しくないです。"
	for _, email := range invalidEmails {
		err := validateEmail(email)
		if err == nil {
			t.Errorf("validateEmail(%q) expected error, but got nil", email)
			continue
		}
		if err.Error() != expectedMsg {
			t.Errorf("validateEmail(%q) error message = %q, want %q",
				email, err.Error(), expectedMsg)
		}
	}
}

func TestValidEmailsDoNotReturnError(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"user@gmail.com",
		"admin@company.co.jp",
		"user.name@example.com",
		"user+tag@example.com",
		"user_name@example.com",
		"user123@example123.com",
		"a@b.co",
		"very.long.username@very.long.domain.name.com",
	}

	for _, email := range validEmails {
		err := validateEmail(email)
		if err != nil {
			t.Errorf("validateEmail(%q) unexpected error: %v", email, err)
		}
	}
}

func BenchmarkNewEmail(b *testing.B) {
	email := "test@example.com"
	for i := 0; i < b.N; i++ {
		_, _ = NewEmail(email)
	}
}

func BenchmarkValidateEmail(b *testing.B) {
	email := "test@example.com"
	for i := 0; i < b.N; i++ {
		_ = validateEmail(email)
	}
}

func TestEmailEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"非常に長いローカル部", "verylongusernamefortestingpurposesonly@example.com", false},
		{"非常に長いドメイン部", "test@verylongdomainnamefortestingpurposesonly.com", false},
		{"最小のメールアドレス", "a@b.co", false},
		{"TLD無し", "test@localhost", true},
		{"特殊文字多数", "user!#$%&@example.com", false},
		{"数字のみローカル部", "123@example.com", false},
		{"数字のみドメイン部", "test@123.com", false},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewEmail(tc.input)
			if tc.wantError {
				if err == nil {
					t.Errorf("NewEmail(%q) expected error, but got nil", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("NewEmail(%q) unexpected error: %v", tc.input, err)
				}
			}
		})
	}
}

func TestNewEmailReturnsPointer(t *testing.T) {
	email1, err1 := NewEmail("test1@example.com")
	email2, err2 := NewEmail("test2@example.com")

	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected errors: %v, %v", err1, err2)
	}

	if email1 == email2 {
		t.Error("NewEmail should return different pointers for different calls")
	}
}
