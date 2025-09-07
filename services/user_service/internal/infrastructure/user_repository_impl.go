package infrastructure

import (
	"context"
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

func (ur UserRepositoryImpl) GetById(ctx context.Context, userId user.UserId) (*aggregate.User, error) {

	return nil, nil
}

func (ur UserRepositoryImpl) Register(ctx context.Context, user *aggregate.User) error {
	tx, err := ur.db.BeginTx(ctx, nil)
	ur.logger.InfoContext(ctx, "トランザクションを開始します。")
	if err != nil {
		return err
	}
	if _, err := tx.Exec(
		"insert into users (user_id, uuid_user_id, email, role_id) values(?, ?, ?, ?)", user.GetUserId().GetValue(), user.GetUuidUserId().GetValue(), user.GetEmail().GetValue(), user.GetRole().GetRoleId()); err != nil {
		ur.logger.InfoContext(ctx, "エラーが発生したためロールバックします。")
		tx.Rollback()
		return err
	}
	if _, err := tx.Exec("insert into user_profile (user_id, user_name, bio) values(?, ?, ?)", user.GetUserId().GetValue(), user.GetUserProfile().GetUserName().GetValue(), user.GetUserProfile().GetBio().GetValue()); err != nil {
		ur.logger.InfoContext(ctx, "エラーが発生したためロールバックします。")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur UserRepositoryImpl) Update(ctx context.Context, user *aggregate.User) error {
	return nil
}

func (ur UserRepositoryImpl) ExistsWithEmail(ctx context.Context, email *user.Email) (bool, error) {
	ur.logger.InfoContext(ctx, "メールアドレスが重複していないか判定します。")
	var _email string
	if err := ur.db.QueryRow("select email from users where email = ?", email.GetValue()).Scan(&_email); err != nil {
		return false, err
	}
	if len(_email) == 0 {
		return true, nil
	}
	return false, nil
}
