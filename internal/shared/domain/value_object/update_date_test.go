package value_object_test

import (
	"testing"
	"time"

	"github.com/ramiromd/crowfunding/internal/shared/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUpdateDate(t *testing.T) {
	t.Run("should create update date when called", func(t *testing.T) {
		before := time.Now()
		ud := value_object.NewUpdateDate()
		after := time.Now()

		assert.False(t, ud.Value() == "")
		assert.False(t, ud.Time().Before(before))
		assert.False(t, ud.Time().After(after))
	})
}

func TestNewUpdateDateFromString(t *testing.T) {
	t.Run("should create update date when string is valid RFC3339", func(t *testing.T) {
		raw := "2024-06-01T12:00:00Z"
		ud, err := value_object.NewUpdateDateFromString(raw)
		require.NoError(t, err)
		assert.Equal(t, raw, ud.Value())
	})

	t.Run("should return error when string is not RFC3339", func(t *testing.T) {
		_, err := value_object.NewUpdateDateFromString("01-06-2024")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidUpdateDateFormat)
	})

	t.Run("should return error when string is empty", func(t *testing.T) {
		_, err := value_object.NewUpdateDateFromString("")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidUpdateDateFormat)
	})
}

func TestUpdateDateValue(t *testing.T) {
	t.Run("should return RFC3339 formatted string when value is called", func(t *testing.T) {
		raw := "2024-06-01T12:00:00Z"
		ud, err := value_object.NewUpdateDateFromString(raw)
		require.NoError(t, err)
		assert.Equal(t, raw, ud.Value())
	})
}

func TestUpdateDateCheckGreaterOrEqualThan(t *testing.T) {
	t.Run("should return nil when update date is after creation date", func(t *testing.T) {
		cd, err := value_object.NewCreationDateFromString("2024-01-01T00:00:00Z")
		require.NoError(t, err)

		ud, err := value_object.NewUpdateDateFromString("2024-06-01T00:00:00Z")
		require.NoError(t, err)

		assert.NoError(t, ud.CheckGreaterOrEqualThan(cd))
	})

	t.Run("should return nil when update date equals creation date", func(t *testing.T) {
		raw := "2024-01-01T00:00:00Z"
		cd, err := value_object.NewCreationDateFromString(raw)
		require.NoError(t, err)

		ud, err := value_object.NewUpdateDateFromString(raw)
		require.NoError(t, err)

		assert.NoError(t, ud.CheckGreaterOrEqualThan(cd))
	})

	t.Run("should return error when update date is before creation date", func(t *testing.T) {
		cd, err := value_object.NewCreationDateFromString("2024-06-01T00:00:00Z")
		require.NoError(t, err)

		ud, err := value_object.NewUpdateDateFromString("2024-01-01T00:00:00Z")
		require.NoError(t, err)

		err = ud.CheckGreaterOrEqualThan(cd)
		assert.ErrorIs(t, err, value_object.ErrUpdateDateBeforeCreationDate)
	})
}

// mustParseRFC3339 parses a RFC3339 string in test helpers, failing the test on error.
func mustParseRFC3339(t *testing.T, s string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, s)
	require.NoError(t, err)
	return parsed
}
