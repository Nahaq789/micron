package services

import (
	"errors"
	"user_service/internal/domain/models/user"
	"user_service/internal/domain/repositories"
)

type EmailDuplicateChekcker interface {
	CheckDuplicate(email *user.Email) error
}

type emailDupulicateService struct {
	repository repositories.UserRepository
}

func NewEmailDuplicateService(r repositories.UserRepository) EmailDuplicateChekcker {
	return &emailDupulicateService{repository: r}
}

func (s *emailDupulicateService) CheckDuplicate(email *user.Email) error {
	exist, err := s.repository.ExistsWithEmail(email)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("このメールアドレスはすでに使用されています。")
	}
	return nil
}
