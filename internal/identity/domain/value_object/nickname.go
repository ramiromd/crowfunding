package value_object

import(
	"errors"
	"regexp"
)

var nicknameRegex = regexp.MustCompile(`^[a-z0-9_]{8,12}$`)
var ErrInvalidNicknameFormat = errors.New("invalid nickname format")
type Nickname struct {
	value string
}

func NewNickname(nickname string) (Nickname, error) {
	if !nicknameRegex.MatchString(nickname) {
		return Nickname{}, ErrInvalidNicknameFormat
	}
	return Nickname{value: nickname}, nil
}

func (n Nickname) Value() string {
	return n.value
}