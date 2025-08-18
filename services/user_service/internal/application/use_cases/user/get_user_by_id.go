package user

import (
	"context"
	"log/slog"
	"user_service/internal/application/dtos"
	"user_service/internal/domain/models/user"
	"user_service/internal/domain/repositories"
)

type GetUserById struct {
	logger     *slog.Logger
	repository repositories.UserRepository
}

func NewGetUserById(l *slog.Logger, repository repositories.UserRepository) *GetUserById {
	return &GetUserById{
		logger:     l,
		repository: repository,
	}
}

func (g *GetUserById) Execute(ctx context.Context, id int) (*dtos.UserDto, error) {
	userId := user.NewUserId(id)
	user, err := g.repository.GetById(ctx, userId)
	if err != nil {
		g.logger.ErrorContext(ctx, "ユーザ情報の取得に失敗しました。", "error", err)
		return nil, err
	}

	dto := dtos.NewUserDto(user.GetUserId().GetValue(), user.GetUuidUserId().GetValue(), user.GetUserProfile().GetUserName().GetValue(), user.GetEmail().GetValue(), user.GetRole().GetRoleId(), user.GetUserType().GetTypeId(), user.GetUserProfile().GetBio().GetValue())
	return dto, nil
}
