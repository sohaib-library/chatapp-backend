package auth

import (
	"errors"

	"chatapp-backend/models"
	"context"
)

var (
	ErrInvalidRequest = errors.New("invalid signup request")
	ErrInvalidEmail   = errors.New("enter valid email")
	ErrUserExists     = errors.New("user already exists")
)

type AuthService interface {
	SignUp(ctx context.Context, user models.Users) error
}
