package aggregate

import (
	"testing"
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
)

func TestNewUserProfile(t *testing.T) {
	tests := []struct {
		name          string
		userProfileId int
		userId        int
		userNameValue string
		bioValue      string
	}{
		{
			name:          "valid user profile",
			userProfileId: 1,
			userId:        100,
			userNameValue: "testuser",
			bioValue:      "I am a test user",
		},
		{
			name:          "user profile with empty bio",
			userProfileId: 2,
			userId:        200,
			userNameValue: "user2",
			bioValue:      "",
		},
		{
			name:          "user profile with long bio",
			userProfileId: 3,
			userId:        300,
			userNameValue: "longbiouser",
			bioValue:      "This is a very long bio that contains detailed information about the user",
		},
		{
			name:          "user profile with zero ids",
			userProfileId: 0,
			userId:        0,
			userNameValue: "zerouser",
			bioValue:      "Zero ID user",
		},
		{
			name:          "user profile with negative ids",
			userProfileId: -1,
			userId:        -5,
			userNameValue: "negativeuser",
			bioValue:      "Negative ID user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userName, err := userprofile.NewUserName(tt.userNameValue)
			if err != nil {
				t.Fatalf("Failed to create UserName: %v", err)
			}

			bio := userprofile.NewBio(tt.bioValue)

			userId := user.NewUserId(tt.userId)
			result := NewUserProfile(tt.userProfileId, userId, userName, bio)

			if result.userProfileId != tt.userProfileId {
				t.Errorf("NewUserProfile() userProfileId = %v, want %v", result.userProfileId, tt.userProfileId)
			}
			if result.userId != userId {
				t.Errorf("NewUserProfile() userId = %v, want %v", result.userId, tt.userId)
			}
			if result.userName != userName {
				t.Errorf("NewUserProfile() userName = %v, want %v", result.userName, userName)
			}
			if result.bio != bio {
				t.Errorf("NewUserProfile() bio = %v, want %v", result.bio, bio)
			}
		})
	}
}

func TestNewUserProfileWithDefaults(t *testing.T) {
	tests := []struct {
		name          string
		userNameValue string
		bioValue      string
	}{
		{
			name:          "valid user profile with defaults",
			userNameValue: "testuser",
			bioValue:      "I am a test user",
		},
		{
			name:          "user profile with empty bio",
			userNameValue: "user2",
			bioValue:      "",
		},
		{
			name:          "user profile with long bio",
			userNameValue: "longbiouser",
			bioValue:      "This is a very long bio that contains detailed information about the user",
		},
		{
			name:          "user profile with zero user id",
			userNameValue: "zerouser",
			bioValue:      "Zero ID user",
		},
		{
			name:          "user profile with negative user id",
			userNameValue: "negativeuser",
			bioValue:      "Negative ID user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userName, err := userprofile.NewUserName(tt.userNameValue)
			if err != nil {
				t.Fatalf("Failed to create UserName: %v", err)
			}

			bio := userprofile.NewBio(tt.bioValue)

			result := NewUserProfileWithDefaults(userName, bio)

			if result.userProfileId != 0 {
				t.Errorf("NewUserProfileWithDefaults() userProfileId = %v, want 0", result.userProfileId)
			}
			if result.userName != userName {
				t.Errorf("NewUserProfileWithDefaults() userName = %v, want %v", result.userName, userName)
			}
			if result.bio != bio {
				t.Errorf("NewUserProfileWithDefaults() bio = %v, want %v", result.bio, bio)
			}
		})
	}
}
