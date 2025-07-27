package infrastructure

import (
	"database/sql"
	"log/slog"
	aggregate "user_service/internal/domain/aggregates"
	"user_service/internal/domain/models/user"
)

type UserRepositoryImpl struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewUserRepositoryImpl(l *slog.Logger, db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		logger: l,
		db:     db,
	}
}

func (ur UserRepositoryImpl) GetById(userId user.UserId) (*aggregate.User, error) {
	return nil, nil
}

func (ur UserRepositoryImpl) Register(user *aggregate.User) error {
	return nil
}

func (ur UserRepositoryImpl) Update(user *aggregate.User) error {
	return nil
}

func (ur UserRepositoryImpl) ExistsWithEmail(email *user.Email) (bool, error) {
	return false, nil
}
