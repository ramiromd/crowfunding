package value_object

import (
	"errors"
	"time"
)

var ErrInvalidUpdateDateFormat = errors.New("invalid update date format")
var ErrUpdateDateBeforeCreationDate = errors.New("update date cannot be before creation date")

type UpdateDate struct {
	value time.Time
}

func NewUpdateDate() UpdateDate {
	return UpdateDate{value: time.Now()}
}

func NewUpdateDateFromString(updateDate string) (UpdateDate, error) {
	parsedUpdateDate, err := time.Parse(time.RFC3339Nano, updateDate)
	if err != nil {
		return UpdateDate{}, ErrInvalidUpdateDateFormat
	}
	return UpdateDate{value: parsedUpdateDate}, nil
}

func (u UpdateDate) Value() string {
	return u.value.Format(time.RFC3339Nano)
}

// TODO: Implement better way to compare dates without relying on string parsing in tests
func (u UpdateDate) Time() time.Time {
	return u.value
}

func (u UpdateDate) CheckGreaterOrEqualThan(creationDate CreationDate) error {
	if u.value.Before(creationDate.Time()) {
		return ErrUpdateDateBeforeCreationDate
	}
	return nil
}