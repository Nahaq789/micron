package dtos

import "github.com/google/uuid"

type UserDto struct {
	userId     int
	uuidUserId uuid.UUID
	email      string
	userName   string
	bio        string
}

func NewUserDto(userId int, uuidUserId uuid.UUID, email, userName, bio string) UserDto {
	return UserDto{
		userId:     userId,
		uuidUserId: uuidUserId,
		email:      email,
		userName:   userName,
		bio:        bio,
	}
}

func (u UserDto) GetUserId() int {
	return u.userId
}

func (u UserDto) GetUuidUserId() uuid.UUID {
	return u.uuidUserId
}

func (u UserDto) GetEmail() string {
	return u.email
}

func (u UserDto) GetUserName() string {
	return u.userName
}

func (u UserDto) GetBio() string {
	return u.bio
}
