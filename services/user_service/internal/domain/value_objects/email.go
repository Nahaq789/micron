package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(v string) (*Email, error) {
	if err := validateEmail(v); err != nil {
		return nil, err
	}

	return &Email{value: v}, nil
}

func validateEmail(value string) error {
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)

	if re.MatchString(value) {
		return errors.New("メールアドレスの形式が正しくないです。")
	}

	return nil
}
