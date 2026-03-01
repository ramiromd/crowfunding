package value_object_test

import (
	"testing"

	"github.com/ramiromd/crowfunding/internal/shared/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEntityId(t *testing.T) {
	t.Run("should create entity id when uuid is valid", func(t *testing.T) {
		validUUID := "123e4567-e89b-12d3-a456-426614174000"
		entityId, err := value_object.NewEntityId(validUUID)
		require.NoError(t, err)
		assert.Equal(t, validUUID, entityId.Value())
	})

	t.Run("should return error when uuid is invalid", func(t *testing.T) {
		_, err := value_object.NewEntityId("invalid-uuid")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidEntityIdFormat)
	})
}
