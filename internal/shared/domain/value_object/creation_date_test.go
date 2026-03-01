package value_object_test

import (
	"testing"
	"time"

	"github.com/ramiromd/crowfunding/internal/shared/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCreationDate(t *testing.T) {
	t.Run("should create creation date when called", func(t *testing.T) {
		before := time.Now()
		cd := value_object.NewCreationDate()
		after := time.Now()
		assert.False(t, cd.Time().Before(before))
		assert.False(t, cd.Time().After(after))
	})
}

func TestNewCreationDateFromString(t *testing.T) {
	t.Run("should create creation date when string is valid RFC3339", func(t *testing.T) {
		raw := "2024-01-15T10:30:00Z"
		cd, err := value_object.NewCreationDateFromString(raw)
		require.NoError(t, err)
		assert.Equal(t, raw, cd.Value())
	})

	t.Run("should return error when string is not RFC3339", func(t *testing.T) {
		_, err := value_object.NewCreationDateFromString("15-01-2024")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidCreationDateFormat)
	})

	t.Run("should return error when string is empty", func(t *testing.T) {
		_, err := value_object.NewCreationDateFromString("")
		require.Error(t, err)
		assert.ErrorIs(t, err, value_object.ErrInvalidCreationDateFormat)
	})
}

func TestCreationDateValue(t *testing.T) {
	t.Run("should return RFC3339 formatted string when value is called", func(t *testing.T) {
		raw := "2024-06-01T00:00:00Z"
		cd, err := value_object.NewCreationDateFromString(raw)
		require.NoError(t, err)
		assert.Equal(t, raw, cd.Value())
	})
}

func TestCreationDateTime(t *testing.T) {
	t.Run("should return underlying time when time is called", func(t *testing.T) {
		raw := "2024-06-01T00:00:00Z"
		expected, _ := time.Parse(time.RFC3339, raw)
		cd, err := value_object.NewCreationDateFromString(raw)
		require.NoError(t, err)
		assert.Equal(t, expected, cd.Time())
	})
}
