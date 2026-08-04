package auth

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, user models.Users) error
	UserExists(ctx context.Context, email string) (bool, error)
	Login(ctx context.Context, email string) (*models.Users, error)
}
