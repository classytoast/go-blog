package users_service

import (
	"context"
	"fmt"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_errors "github.com/classytoast/go-blog/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) RegisterUser(ctx context.Context, user *domain.NewUserUninitialized) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("validate user domain: %w", err)
	}

	exists, err := s.repo.UserExists(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("check user exists: %w", err)
	}
	if exists {
		return fmt.Errorf("check user exists: %w", core_errors.ErrUserAlreadyExists)
	}

	hashPass, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("hash user password: %w", err)
	}

	if err = s.repo.CreateUser(ctx, user, hashPass); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

// HashPassword принимает чистый пароль и возвращает его bcrypt-хеш в виде строки
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
