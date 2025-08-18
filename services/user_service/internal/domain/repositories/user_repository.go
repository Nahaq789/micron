package repositories

import (
	"context"
	aggregate "user_service/internal/domain/aggregates"
	"user_service/internal/domain/models/user"
)

type UserRepository interface {
	GetById(ctx context.Context, userId user.UserId) (*aggregate.User, error)
	Register(ctx context.Context, user *aggregate.User) error
	Update(ctx context.Context, user *aggregate.User) error
	ExistsWithEmail(ctx context.Context, email *user.Email) (bool, error)
}
