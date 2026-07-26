package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"chatapp-backend/utils"
	"context"
	"errors"
	"strings"
)

func (s *AuthServiceImpl) Login(ctx context.Context, login models.Login) (string, error) {
	login.Email = strings.TrimSpace(login.Email)
	login.Password = strings.TrimSpace(login.Password)

	if login.Email == "" || login.Password == "" {
		return "", apperror.ErrInvalidCredentials
	}

	if !emailRegex.MatchString(login.Email) {
		return "", apperror.ErrInvalidEmail
	}

	user, err := s.repo.Login(ctx, login.Email)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			return "", apperror.ErrInvalidCredentials
		}
		return "", err
	}

	if !utils.CheckPassword(login.Password, user.Password) {
		return "", apperror.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
