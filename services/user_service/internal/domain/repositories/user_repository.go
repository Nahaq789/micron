package repositories

import aggregate "user_service/internal/domain/aggregates"

type UserRepository interface {
	GetById(userId int) (*aggregate.User, error)
	Register(user *aggregate.User) error
	Update(user *aggregate.User) error
}
