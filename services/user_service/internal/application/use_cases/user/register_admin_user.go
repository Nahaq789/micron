package user

import (
	"context"
	"log/slog"
	"user_service/internal/application/commands"
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
	"user_service/internal/domain/repositories"
)

type RegisterAdminUser struct {
	logger     slog.Logger
	repository repositories.UserRepository
}

func NewCreateUser(l slog.Logger, r repositories.UserRepository) RegisterAdminUser {
	return RegisterAdminUser{logger: l, repository: r}
}

func (r RegisterAdminUser) RegisterAdmin(ctx context.Context, c commands.RegisterAdminUserCommand) error {
	email, err := user.NewEmail(c.GetEmail())
	if err != nil {
		r.logger.ErrorContext(ctx, "メールアドレスが不正です。", "error", err)
	}

	userName, err := userprofile.NewUserName(c.GetUserName())
	if err != nil {
		r.logger.ErrorContext(ctx, "ユーザ名が不正です。", "error", err)
	}

	bio := userprofile.NewBio(c.GetBio())

	return nil
}
