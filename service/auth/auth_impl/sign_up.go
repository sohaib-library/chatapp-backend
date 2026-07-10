package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"context"
	"errors"
	"fmt"
)

func (s *AuthServiceImpl) SignUp(ctx context.Context, user models.Users) error {

	exists, err := s.repo.UserExists(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("check existing user: %w", err)
	}


	if exists {
		return apperror.ErrUserExists
	}
	

	if err := s.repo.CreateUser(ctx, user); err != nil {
		if errors.Is(err, apperror.ErrUserExists) {
			return err
		}
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
