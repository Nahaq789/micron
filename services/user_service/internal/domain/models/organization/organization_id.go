package organization

import (
	"errors"
	"fmt"
	"strings"

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
	value = strings.TrimSpace(value)
	if value == "" {
		return OrganizationId{}, errors.New("organization ID cannot be empty")
	}

	u, err := uuid.Parse(value)
	if err != nil {
		return OrganizationId{}, fmt.Errorf("invalid organization ID format: %w", err)
	}

	// UUID version チェック（必要に応じて）
	if u.Version() != 4 {
		return OrganizationId{}, errors.New("organization ID must be a valid UUID v4")
	}

	return OrganizationId{value: u}, nil
}
