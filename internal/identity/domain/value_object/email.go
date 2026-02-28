package value_object

import (
	"net/mail"
	"errors"
)

var ErrInvalidEmailFormat = errors.New("invalid email format")

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return Email{}, ErrInvalidEmailFormat
	}
	return Email{value: email}, nil
}

func (e Email) Value() string {
	return e.value
}