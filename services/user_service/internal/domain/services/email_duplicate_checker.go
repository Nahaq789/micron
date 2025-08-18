package services

import (
	"context"
	"errors"
	"user_service/internal/domain/models/user"
	"user_service/internal/domain/repositories"
)

type EmailDuplicateChecker interface {
	CheckDuplicate(ctx context.Context, email *user.Email) error
}

type EmailDuplicateService struct {
	repository repositories.UserRepository
}

func NewEmailDuplicateService(r repositories.UserRepository) EmailDuplicateChecker {
	return &EmailDuplicateService{repository: r}
}

func (s *EmailDuplicateService) CheckDuplicate(ctx context.Context, email *user.Email) error {
	exist, err := s.repository.ExistsWithEmail(ctx, email)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("このメールアドレスはすでに使用されています。")
	}
	return nil
}
