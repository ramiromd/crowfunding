package value_object

import (
	"errors"
	"regexp"
)

var bcryptRegex = regexp.MustCompile(`^\$2[ayb]\$`)
var ErrInvalidPasswordHashFormat = errors.New("invalid password hash format")

type PasswordHash struct {
	value string
}

func NewPasswordHash(passwordHash string) (PasswordHash, error) {
	if !bcryptRegex.MatchString(passwordHash) {
		return PasswordHash{}, ErrInvalidPasswordHashFormat
	}
	return PasswordHash{value: passwordHash}, nil
}

func (p PasswordHash) Value() string {
	return p.value
}