package repositories

import (
	aggregate "user_service/internal/domain/aggregates"
	"user_service/internal/domain/models/user"
)

type UserRepository interface {
	GetById(userId int) (*aggregate.User, error)
	Register(user *aggregate.User) error
	Update(user *aggregate.User) error
	ExistsWithEmail(email *user.Email) (bool, error)
}
