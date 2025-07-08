package organization

import (
	"bytes"
	"errors"

	"github.com/google/uuid"
)

type OrganizationId struct {
	value uuid.UUID
}

type UUIDGenerator func() (uuid.UUID, error)

var defaultUUIDGenerator UUIDGenerator = uuid.NewV7

func NewOrganizationId() (OrganizationId, error) {
	return NewOrganizationIdWithGenerator(defaultUUIDGenerator)
}

func NewOrganizationIdWithGenerator(generator UUIDGenerator) (OrganizationId, error) {
	u, err := generator()
	if err != nil {
		return OrganizationId{}, errors.New("UUID v7の生成に失敗しました。")
	}

	return OrganizationId{value: u}, nil
}

func FromOrganizationId(value string) (OrganizationId, error) {
	if len(value) == 0 {
		return OrganizationId{}, errors.New("文字列が空です。")
	}
	b := bytes.NewBufferString(value)
	u, err := uuid.FromBytes(b.Bytes())
	if err != nil {
		return OrganizationId{}, errors.New("文字列からUUIDの変換に失敗しました。")
	}
	return OrganizationId{value: u}, nil
}
