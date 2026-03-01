package value_object

import (
	"errors"
	"time"
)

var ErrInvalidCreationDateFormat = errors.New("invalid creation date format")

type CreationDate struct {
	value time.Time
}

func NewCreationDate() CreationDate {
	return CreationDate{value: time.Now()}
}

func NewCreationDateFromString(creationDate string) (CreationDate, error) {
	parsedCreationDate, err := time.Parse(time.RFC3339Nano, creationDate)
	if err != nil {
		return CreationDate{}, ErrInvalidCreationDateFormat
	}
	return CreationDate{value: parsedCreationDate}, nil
}

func (c CreationDate) Value() string {
	return c.value.Format(time.RFC3339Nano)
}

func (c CreationDate) Time() time.Time {
	return c.value
}