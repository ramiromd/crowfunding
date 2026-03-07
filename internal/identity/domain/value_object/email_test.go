package value_object_test

import (
	"testing"

	"github.com/ramiromd/crowfunding/internal/identity/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmail(t *testing.T) {
	t.Run("should create email when format is valid", func(t *testing.T) {
		email, err := value_object.NewEmail("user@example.com")
		require.NoError(t, err)
		assert.Equal(t, "user@example.com", email.Value())
	})

	t.Run("should return error when email has no at sign", func(t *testing.T) {
		_, err := value_object.NewEmail("userexample.com")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidEmailFormat)
	})

	t.Run("should return error when email has no domain", func(t *testing.T) {
		_, err := value_object.NewEmail("user@")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidEmailFormat)
	})

	t.Run("should return error when email is empty", func(t *testing.T) {
		_, err := value_object.NewEmail("")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidEmailFormat)
	})

	t.Run("should return error when email has no local part", func(t *testing.T) {
		_, err := value_object.NewEmail("@example.com")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidEmailFormat)
	})
}
