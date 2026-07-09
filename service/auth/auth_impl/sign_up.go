package auth_impl

import (
	"chatapp-backend/models"
	authService "chatapp-backend/service/auth"
	"context"
	"fmt"
	"strings"
)

func (s *AuthServiceImpl) SignUp(ctx context.Context, user models.Users) error {
	if strings.TrimSpace(user.NAME) == "" ||
		strings.TrimSpace(user.EMAIL) == "" ||
		strings.TrimSpace(user.PASSWORD) == "" {
		return authService.ErrInvalidRequest
	}

	exists, err := s.repo.UserExists(ctx, user.NAME, user.EMAIL)
	if err != nil {
		return fmt.Errorf("check existing user: %w", err)
	}
	if exists {
		return authService.ErrUserExists
	}

	if err := s.repo.CreateUser(ctx, user.NAME, user.EMAIL, user.PASSWORD); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil

}
