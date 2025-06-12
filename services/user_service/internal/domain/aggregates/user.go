package aggregate

import (
	"user_service/internal/domain/models/user"
	usertype "user_service/internal/domain/models/user_type"
)

type User struct {
	userId      int
	uuidUserId  user.UUIDUserId
	email       user.Email
	rollId      int
	userType    usertype.UserType
	userProfile UserProfile
}

func NewUser(userId int, uuidUserId user.UUIDUserId, email user.Email, rollId int, userType usertype.UserType, userProfile UserProfile) *User {
	return &User{
		userId:      userId,
		uuidUserId:  uuidUserId,
		email:       email,
		rollId:      rollId,
		userType:    userType,
		userProfile: userProfile,
	}
}

// データ登録用
func NewUserWithDefaults(uuidUserId user.UUIDUserId, email user.Email, rollId int, userType usertype.UserType, userProfile UserProfile) *User {
	return &User{
		userId:      0,
		uuidUserId:  uuidUserId,
		email:       email,
		rollId:      rollId,
		userType:    userType,
		userProfile: userProfile,
	}
}
