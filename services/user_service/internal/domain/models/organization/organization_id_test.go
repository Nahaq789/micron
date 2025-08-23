package organization

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrganizationId(t *testing.T) {
	orgId, err := NewOrganizationId()
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, orgId.value)
	assert.Equal(t, uuid.Version(7), orgId.value.Version())
}

func TestNewOrganizationIdWithGenerator(t *testing.T) {
	t.Run("success with UUID v7", func(t *testing.T) {
		generator := func() (uuid.UUID, error) {
			return uuid.NewV7()
		}

		orgId, err := NewOrganizationIdWithGenerator(generator)
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, orgId.value)
	})

	t.Run("success with custom UUID", func(t *testing.T) {
		expectedUUID := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		generator := func() (uuid.UUID, error) {
			return expectedUUID, nil
		}

		orgId, err := NewOrganizationIdWithGenerator(generator)
		assert.NoError(t, err)
		assert.Equal(t, expectedUUID, orgId.value)
	})

	t.Run("generator returns error", func(t *testing.T) {
		generator := func() (uuid.UUID, error) {
			return uuid.Nil, errors.New("generation failed")
		}

		orgId, err := NewOrganizationIdWithGenerator(generator)
		assert.Error(t, err)
		assert.Equal(t, "UUID v7の生成に失敗しました。", err.Error())
		assert.Equal(t, uuid.Nil, orgId.value)
	})
}

func TestFromOrganizationId(t *testing.T) {
	t.Run("success with valid UUID v4", func(t *testing.T) {
		validUUIDv4 := "550e8400-e29b-41d4-a716-446655440000"

		orgId, err := FromOrganizationId(validUUIDv4)
		assert.NoError(t, err)
		assert.Equal(t, uuid.MustParse(validUUIDv4), orgId.value)
	})

	t.Run("success with UUID v4 with whitespace", func(t *testing.T) {
		validUUIDv4 := "  550e8400-e29b-41d4-a716-446655440000  "

		orgId, err := FromOrganizationId(validUUIDv4)
		assert.NoError(t, err)
		assert.Equal(t, uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), orgId.value)
	})

	t.Run("empty string", func(t *testing.T) {
		orgId, err := FromOrganizationId("")
		assert.Error(t, err)
		assert.Equal(t, "organization ID cannot be empty", err.Error())
		assert.Equal(t, uuid.Nil, orgId.value)
	})

	t.Run("whitespace only", func(t *testing.T) {
		orgId, err := FromOrganizationId("   ")
		assert.Error(t, err)
		assert.Equal(t, "organization ID cannot be empty", err.Error())
		assert.Equal(t, uuid.Nil, orgId.value)
	})

	t.Run("invalid UUID format", func(t *testing.T) {
		invalidUUID := "invalid-uuid"

		orgId, err := FromOrganizationId(invalidUUID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid organization ID format")
		assert.Equal(t, uuid.Nil, orgId.value)
	})

	t.Run("valid UUID but not v4", func(t *testing.T) {
		uuidV1 := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"

		orgId, err := FromOrganizationId(uuidV1)
		assert.Error(t, err)
		assert.Equal(t, "organization ID must be a valid UUID v4", err.Error())
		assert.Equal(t, uuid.Nil, orgId.value)
	})

	t.Run("UUID v7", func(t *testing.T) {
		uuidV7, _ := uuid.NewV7()

		orgId, err := FromOrganizationId(uuidV7.String())
		assert.Error(t, err)
		assert.Equal(t, "organization ID must be a valid UUID v4", err.Error())
		assert.Equal(t, uuid.Nil, orgId.value)
	})

	t.Run("nil UUID", func(t *testing.T) {
		nilUUID := "00000000-0000-0000-0000-000000000000"

		orgId, err := FromOrganizationId(nilUUID)
		assert.Error(t, err)
		assert.Equal(t, "organization ID must be a valid UUID v4", err.Error())
		assert.Equal(t, uuid.Nil, orgId.value)
	})
}

func TestOrganizationIdValue(t *testing.T) {
	expectedUUID := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
	orgId := OrganizationId{value: expectedUUID}

	assert.Equal(t, expectedUUID, orgId.value)
}

func BenchmarkNewOrganizationId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewOrganizationId()
	}
}

func BenchmarkFromOrganizationId(b *testing.B) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FromOrganizationId(validUUID)
	}
}
