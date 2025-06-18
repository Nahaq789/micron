package aggregate

import (
	"user_service/internal/domain/models/role"
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
	usertype "user_service/internal/domain/models/user_type"
)

type User struct {
	userId      int
	uuidUserId  user.UUIDUserId
	email       *user.Email
	role        role.Role
	userType    usertype.UserType
	userProfile UserProfile
}

func (u User) UpdateUserProfile(user *User, userName userprofile.UserName, bio userprofile.Bio) *User {
	newProfile := NewUserProfile(user.userProfile.userId, user.userId, userName, bio)

	return &User{
		userId:      user.userId,
		uuidUserId:  user.uuidUserId,
		email:       user.email,
		role:        user.role,
		userType:    user.userType,
		userProfile: newProfile,
	}
}

func RegisterAdminUser(email *user.Email, userName userprofile.UserName, bio userprofile.Bio) (*User, error) {
	userProfile := NewUserProfileWithDefaults(userName, bio)

	admin := role.DetermineAdminRole()
	member := usertype.NewMember()

	uuid, err := user.NewUuidUserId()
	if err != nil {
		return nil, err
	}

	user := &User{
		userId:      0,
		uuidUserId:  uuid,
		email:       email,
		role:        admin,
		userType:    member,
		userProfile: userProfile,
	}

	return user, nil
}
