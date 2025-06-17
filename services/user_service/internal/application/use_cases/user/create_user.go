package user

import (
	"log/slog"
	"user_service/internal/domain/repositories"
)

type CreateUser struct {
	logger     slog.Logger
	repository repositories.UserRepository
}

func NewCreateUser(l slog.Logger, r repositories.UserRepository) CreateUser {
	return CreateUser{logger: l, repository: r}
}
