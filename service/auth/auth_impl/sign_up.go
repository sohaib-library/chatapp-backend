package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"chatapp-backend/utils"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func (s *AuthServiceImpl) SignUp(ctx context.Context, user models.Users) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return apperror.ErrInvalidRequest
	}

	if !emailRegex.MatchString(user.Email) {
		return apperror.ErrInvalidEmail
	}

	exists, err := s.repo.UserExists(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("check existing user: %w", err)
	}

	if exists {
		return apperror.ErrUserExists
	}

	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	user.Password = hashedPassword

	if err := s.repo.CreateUser(ctx, user); err != nil {
		if errors.Is(err, apperror.ErrUserExists) {
			return err
		}
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
