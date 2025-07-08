package user

import (
	"context"
	"log/slog"
	"user_service/internal/application/commands"
	aggregate "user_service/internal/domain/aggregates"
	"user_service/internal/domain/models/organization"
	"user_service/internal/domain/models/user"
	userprofile "user_service/internal/domain/models/user_profile"
	"user_service/internal/domain/repositories"
	"user_service/internal/domain/services"
)

type RegisterAdminUser struct {
	logger     slog.Logger
	checker    services.EmailDuplicateChecker
	repository repositories.UserRepository
}

func NewCreateUser(l slog.Logger, c services.EmailDuplicateChecker, r repositories.UserRepository) *RegisterAdminUser {
	return &RegisterAdminUser{logger: l, checker: c, repository: r}
}

func (r *RegisterAdminUser) RegisterAdmin(ctx context.Context, c commands.RegisterAdminUserCommand) error {
	email, err := user.NewEmail(c.GetEmail())
	if err != nil {
		r.logger.ErrorContext(ctx, "メールアドレスが不正です。", "error", err)
		return err
	}

	if exists := r.checker.CheckDuplicate(email); exists != nil {
		r.logger.ErrorContext(ctx, "メールアドレス重複チェックでエラーが発生しました。", "error", exists)
		return exists
	}

	userName, err := userprofile.NewUserName(c.GetUserName())
	if err != nil {
		r.logger.ErrorContext(ctx, "ユーザ名が不正です。", "error", err)
		return err
	}

	bio := userprofile.NewBio(c.GetBio())
	organizationId, err := organization.FromOrganizationId(c.GetOrganizationId())

	user, err := aggregate.RegisterAdminUser(email, userName, bio, organizationId)
	if err != nil {
		r.logger.ErrorContext(ctx, "管理者ユーザの作成に失敗しました。", "error", err)
		return err
	}

	registerErr := r.repository.Register(user)
	if registerErr != nil {
		r.logger.ErrorContext(ctx, "管理者ユーザの登録に失敗しました。", "error", err)
		return err
	}

	return nil
}
