package value_object_test

import (
	"testing"

	"github.com/ramiromd/crowfunding/internal/identity/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPasswordHash(t *testing.T) {
	t.Run("should create password hash when prefix is $2a$", func(t *testing.T) {
		hash := "$2a$10$EixZaYVK1fsbw1ZfbX3OXePaWxn96p36zLSUgQDDJoFBZXsXaHXpe"
		ph, err := value_object.NewPasswordHash(hash)
		require.NoError(t, err)
		assert.Equal(t, hash, ph.Value())
	})

	t.Run("should create password hash when prefix is $2b$", func(t *testing.T) {
		hash := "$2b$10$EixZaYVK1fsbw1ZfbX3OXePaWxn96p36zLSUgQDDJoFBZXsXaHXpe"
		ph, err := value_object.NewPasswordHash(hash)
		require.NoError(t, err)
		assert.Equal(t, hash, ph.Value())
	})

	t.Run("should create password hash when prefix is $2y$", func(t *testing.T) {
		hash := "$2y$10$EixZaYVK1fsbw1ZfbX3OXePaWxn96p36zLSUgQDDJoFBZXsXaHXpe"
		ph, err := value_object.NewPasswordHash(hash)
		require.NoError(t, err)
		assert.Equal(t, hash, ph.Value())
	})

	t.Run("should return error when hash is a plain text password", func(t *testing.T) {
		_, err := value_object.NewPasswordHash("mysecretpassword")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidPasswordHashFormat)
	})

	t.Run("should return error when hash has an invalid bcrypt prefix", func(t *testing.T) {
		_, err := value_object.NewPasswordHash("$2c$10$EixZaYVK1fsbw1ZfbX3OXe")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidPasswordHashFormat)
	})

	t.Run("should return error when hash is empty", func(t *testing.T) {
		_, err := value_object.NewPasswordHash("")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidPasswordHashFormat)
	})
}
