package services

import (
	"errors"
	"user_service/internal/domain/models/user"
	"user_service/internal/domain/repositories"
)

type EmailDuplicateChecker interface {
	CheckDuplicate(email *user.Email) error
}

type emailDuplicateService struct {
	repository repositories.UserRepository
}

func NewEmailDuplicateService(r repositories.UserRepository) EmailDuplicateChecker {
	return &emailDuplicateService{repository: r}
}

func (s *emailDuplicateService) CheckDuplicate(email *user.Email) error {
	exist, err := s.repository.ExistsWithEmail(email)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("このメールアドレスはすでに使用されています。")
	}
	return nil
}
