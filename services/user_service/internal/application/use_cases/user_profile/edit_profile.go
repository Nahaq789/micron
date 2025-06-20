package userprofile

import (
	"context"
	"log/slog"
	"user_service/internal/application/commands"
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
	"user_service/internal/domain/repositories"
)

type EditProfile struct {
	logger     slog.Logger
	repository repositories.UserRepository
}

func NewEditProfile(r repositories.UserRepository) EditProfile {
	return EditProfile{repository: r}
}

func (e *EditProfile) EditUserProfile(ctx context.Context, command commands.EditProfileCommand) error {
	userId := user.NewUserId(command.GetUserId())
	user, err := e.repository.GetById(userId)
	if err != nil {
		e.logger.ErrorContext(ctx, "ユーザ情報取得に失敗しました。", "error", err)
		return err
	}

	userName, err := userprofile.NewUserName(command.GetUserName())
	if err != nil {
		e.logger.ErrorContext(ctx, "ユーザ名が不正です。", "error", err)
	}
	bio := userprofile.NewBio(command.GetBio())

	new := user.UpdateUserProfile(user, userName, bio)

	err = e.repository.Update(new)
	if err != nil {
		e.logger.ErrorContext(ctx, "ユーザの更新に失敗しました。", "error", err)
	}

	return nil
}
