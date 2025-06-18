package aggregate

import userprofile "user_service/internal/domain/models/user_profile"

type UserProfile struct {
	userProfileId int
	userId        int
	userName      userprofile.UserName
	bio           userprofile.Bio
}

func NewUserProfile(userProfileId int, userId int, userName userprofile.UserName, bio userprofile.Bio) UserProfile {
	return UserProfile{
		userProfileId: userProfileId,
		userId:        userId,
		userName:      userName,
		bio:           bio,
	}
}

// データ登録用
func NewUserProfileWithDefaults(userName userprofile.UserName, bio userprofile.Bio) UserProfile {
	return UserProfile{
		userProfileId: 0,
		userId:        0,
		userName:      userName,
		bio:           bio,
	}
}
