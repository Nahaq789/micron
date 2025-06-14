package userprofile

import (
	"context"
	"log/slog"
	"user_service/internal/application/commands"
	userprofile "user_service/internal/domain/models/user_profile"
	"user_service/internal/domain/repositories"
)

type EditProfile struct {
	repository repositories.UserRepository
}

func NewEditProfile(r repositories.UserRepository) EditProfile {
	return EditProfile{repository: r}
}

func (e EditProfile) EditUserProfile(ctx context.Context, command commands.EditProfileCommand) error {
	user, err := e.repository.GetById(command.GetUserId())
	if err != nil {
		return err
	}

	userName, err := userprofile.NewUserName(command.GetUserName())
	if err != nil {
		slog.ErrorContext(ctx, "ユーザ名が不正です。", "error", err)
	}
	bio := userprofile.NewBio(command.GetBio())

	user.Update(userName, bio)

	err = e.repository.Update(user)
	if err != nil {
		slog.ErrorContext(ctx, "ユーザの更新に失敗しました。", "error", err)
	}

	return nil
}
