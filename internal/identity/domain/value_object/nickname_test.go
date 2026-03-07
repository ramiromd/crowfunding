package value_object_test

import (
	"testing"

	"github.com/ramiromd/crowfunding/internal/identity/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNickname(t *testing.T) {
	t.Run("should create nickname when format is valid", func(t *testing.T) {
		nickname, err := value_object.NewNickname("user_123")
		require.NoError(t, err)
		assert.Equal(t, "user_123", nickname.Value())
	})

	t.Run("should return error when nickname is too short", func(t *testing.T) {
		_, err := value_object.NewNickname("usr")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidNicknameFormat)
	})

	t.Run("should return error when nickname is too long", func(t *testing.T) {
		_, err := value_object.NewNickname("this_is_a_very_long_nickname")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidNicknameFormat)
	})

	t.Run("should return error when nickname contains uppercase letters", func(t *testing.T) {
		_, err := value_object.NewNickname("UserName123")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidNicknameFormat)
	})

	t.Run("should return error when nickname contains special characters", func(t *testing.T) {
		_, err := value_object.NewNickname("user@123!")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidNicknameFormat)
	})

	t.Run("should return error when nickname is empty", func(t *testing.T) {
		_, err := value_object.NewNickname("")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidNicknameFormat)
	})

	t.Run("should create nickname when length is exactly 8 characters", func(t *testing.T) {
		nickname, err := value_object.NewNickname("user1234")
		require.NoError(t, err)
		assert.Equal(t, "user1234", nickname.Value())
	})

	t.Run("should create nickname when length is exactly 12 characters", func(t *testing.T) {
		nickname, err := value_object.NewNickname("user12345678")
		require.NoError(t, err)
		assert.Equal(t, "user12345678", nickname.Value())
	})
}
