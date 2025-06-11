package user

import (
	"errors"

	"github.com/google/uuid"
)

type UUIDUserId struct {
	value uuid.UUID
}

type UUIDGenerator func() (uuid.UUID, error)

var defaultUUIDGenerator UUIDGenerator = uuid.NewV7

func NewUuidUserId() (UUIDUserId, error) {
	return NewUuidUserIdWithGenerator(defaultUUIDGenerator)
}

func NewUuidUserIdWithGenerator(generator UUIDGenerator) (UUIDUserId, error) {
	u, err := generator()
	if err != nil {
		return UUIDUserId{}, errors.New("UUID v7の生成に失敗しました。")
	}

	return UUIDUserId{value: u}, nil
}
