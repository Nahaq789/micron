package aggregate

import (
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
	usertype "user_service/internal/domain/models/user_type"
)

type User struct {
	userId      int
	uuidUserId  user.UUIDUserId
	email       *user.Email
	rollId      int
	userType    usertype.UserType
	userProfile UserProfile
}

func NewUser(userId int, uuidUserId user.UUIDUserId, email *user.Email, userType usertype.UserType, userProfile UserProfile) *User {
	return &User{
		userId:      userId,
		uuidUserId:  uuidUserId,
		email:       email,
		rollId:      userType.DecideRole(),
		userType:    userType,
		userProfile: userProfile,
	}
}

// データ登録用
func NewUserWithDefaults(uuidUserId user.UUIDUserId, email *user.Email, userType usertype.UserType, userProfile UserProfile) *User {
	return &User{
		userId:      0,
		uuidUserId:  uuidUserId,
		email:       email,
		rollId:      userType.DecideRole(),
		userType:    userType,
		userProfile: userProfile,
	}
}

func (u User) Update(userName userprofile.UserName, bio userprofile.Bio) {
	u.userProfile.userName = userName
	u.userProfile.bio = bio
}
