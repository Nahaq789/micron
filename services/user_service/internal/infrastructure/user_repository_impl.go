package infrastructure

import (
	"database/sql"
	"log/slog"
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
