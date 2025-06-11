package userprofile

import "errors"

type UserName struct {
	value string
}

func NewUserName(v string) (UserName, error) {
	err := validateUserName(v)
	if err != nil {
		return UserName{}, err
	}
	return UserName{value: v}, nil
}

func validateUserName(v string) error {
	if len(v) <= 0 {
		return errors.New("ユーザ名は1文字以上にしてください。")
	}
	return nil
}
