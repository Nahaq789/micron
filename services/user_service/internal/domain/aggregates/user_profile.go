package aggregate

import (
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
)

type UserProfile struct {
	userProfileId int
	userId        user.UserId
	userName      userprofile.UserName
	bio           userprofile.Bio
}

func NewUserProfile(userProfileId int, userId user.UserId, userName userprofile.UserName, bio userprofile.Bio) UserProfile {
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
		userId:        user.Init(),
		userName:      userName,
		bio:           bio,
	}
}

func (u UserProfile) GetUserName() userprofile.UserName {
	return u.userName
}
func (u UserProfile) GetBio() userprofile.Bio {
	return u.bio
}
