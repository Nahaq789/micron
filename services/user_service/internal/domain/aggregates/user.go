package aggregate

import (
	"user_service/internal/domain/models/organization"
	"user_service/internal/domain/models/role"
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
	usertype "user_service/internal/domain/models/user_type"
)

type User struct {
	userId          user.UserId
	uuidUserId      user.UUIDUserId
	email           *user.Email
	role            role.Role
	userType        usertype.UserType
	userProfile     UserProfile
	organization_id organization.OrganizationId
}

func (u User) UpdateUserProfile(user *User, userName userprofile.UserName, bio userprofile.Bio) *User {
	newProfile := NewUserProfile(user.userProfile.userProfileId, user.userId, userName, bio)

	return &User{
		userId:      user.userId,
		uuidUserId:  user.uuidUserId,
		email:       user.email,
		role:        user.role,
		userType:    user.userType,
		userProfile: newProfile,
	}
}

func RegisterAdminUser(email *user.Email,
	userName userprofile.UserName,
	bio userprofile.Bio,
	organizationId organization.OrganizationId) (*User, error) {
	userProfile := NewUserProfileWithDefaults(userName, bio)

	admin := role.DetermineAdminRole()
	member := usertype.NewMember()

	uuid, err := user.NewUuidUserId()
	if err != nil {
		return nil, err
	}

	user := &User{
		userId:          user.Init(),
		uuidUserId:      uuid,
		email:           email,
		role:            admin,
		userType:        member,
		userProfile:     userProfile,
		organization_id: organizationId,
	}

	return user, nil
}

func (u User) GetUserId() user.UserId {
	return u.userId
}
func (u User) GetUuidUserId() user.UUIDUserId {
	return u.uuidUserId
}
func (u User) GetEmail() *user.Email {
	return u.email
}
func (u User) GetRole() role.Role {
	return u.role
}
func (u User) GetUserType() usertype.UserType {
	return u.userType
}
func (u User) GetUserProfile() UserProfile {
	return u.userProfile
}
