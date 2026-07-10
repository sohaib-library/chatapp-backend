package auth

import (
	"chatapp-backend/models"
	"context"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, user models.Users) error
	UserExists(ctx context.Context, email string) (bool, error)
}
