package userprofile

import (
	"user_service/internal/application/dtos"
	"user_service/internal/domain/repositories"
)

type EditProfile struct {
	repository repositories.UserRepository
}

func NewEditProfile(r repositories.UserRepository) EditProfile {
	return EditProfile{repository: r}
}

func (e EditProfile) EditUserProfile(userDto dtos.UserDto) error {
	user, err := e.repository.GetById(userDto.GetUserId())
	if err != nil {
		return err
	}

	user.Update(userDto.GetEmail(), userDto.GetUserName(), userDto.GetBio())
}
