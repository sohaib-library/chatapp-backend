package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"chatapp-backend/utils"
	"context"
)

func (s *AuthServiceImpl) Login(ctx context.Context, login models.Login) (string, error) {

	user, err := s.repo.Login(ctx, login.Email)
	if err != nil {
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
