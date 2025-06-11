package user

import "testing"

func TestUserType_String(t *testing.T) {
	tests := []struct {
		userType UserType
		want     string
	}{
		{Member, "Member"},
		{Guest, "Guest"},
	}

	for _, test := range tests {
		result := test.userType.String()
		if result != test.want {
			t.Errorf("Expected %s.String() to return: %s, but got: %s",
				test.userType.String(), test.want, result)
		}
	}
}

func TestUserType_Int(t *testing.T) {
	tests := []struct {
		userType UserType
		want     int
	}{
		{Member, 1},
		{Guest, 2},
	}

	for _, test := range tests {
		result := test.userType.Int()
		if result != test.want {
			t.Errorf("Expected %s to have value: %d, but got: %d",
				test.userType, test.want, result)
		}
	}
}

func TestUserTypeStringForUndefined(t *testing.T) {
	undefinedType := UserType(99)
	result := undefinedType.String()
	if result != "" {
		t.Errorf("Expected undefined UserType to return empty string, but got '%s'", result)
	}
}

func TestUserTypeName(t *testing.T) {
	tests := []struct {
		userType UserType
		expected string
	}{
		{Member, "Member"},
		{Guest, "Guest"},
	}

	for _, test := range tests {
		if name, exists := userTypeName[test.userType]; !exists {
			t.Errorf("userTypeName should contain key %d", int(test.userType))
		} else if name != test.expected {
			t.Errorf("Expected userTypeName[%d] to be '%s', but got '%s'",
				int(test.userType), test.expected, name)
		}
	}
}

func TestUserTypeKey(t *testing.T) {
	tests := []struct {
		userType UserType
		expected int
	}{
		{Member, 1},
		{Guest, 2},
	}

	for _, test := range tests {
		if key, exists := userTypeKey[test.userType]; !exists {
			t.Errorf("userTypeKey should contain key %d", int(test.userType))
		} else if key != test.expected {
			t.Errorf("Expected userTypeKey[%s] to be %d, but got %d",
				test.userType.String(), test.expected, key)
		}
	}
}

func TestUserTypeMapConsistency(t *testing.T) {
	if len(userTypeName) != len(userTypeKey) {
		t.Errorf("userTypeName and userTypeKey should have the same number of entries")
	}

	for userType := range userTypeName {
		if _, exists := userTypeKey[userType]; !exists {
			t.Errorf("userTypeKey is missing entry for %s", userType.String())
		}
	}

	for userType := range userTypeKey {
		if _, exists := userTypeName[userType]; !exists {
			t.Errorf("userTypeName is missing entry for UserType %d", int(userType))
		}
	}
}

// ベンチマークテスト
func BenchmarkUserTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Member.String()
	}
}

func TestAllUserTypes(t *testing.T) {
	testCases := []struct {
		name     string
		userType UserType
		wantName string
		wantKey  int
	}{
		{
			name:     "Member",
			userType: Member,
			wantName: "Member",
			wantKey:  1,
		},
		{
			name:     "Guest",
			userType: Guest,
			wantName: "Guest",
			wantKey:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.userType.String(); got != tc.wantName {
				t.Errorf("String() = %q, want %q", got, tc.wantName)
			}

			if got := userTypeName[tc.userType]; got != tc.wantName {
				t.Errorf("userTypeName[%v] = %q, want %q", tc.userType, got, tc.wantName)
			}

			if got := userTypeKey[tc.userType]; got != tc.wantKey {
				t.Errorf("userTypeKey[%v] = %d, want %d", tc.userType, got, tc.wantKey)
			}
		})
	}
}
