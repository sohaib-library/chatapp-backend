package auth

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

type AuthService interface {
	SignUp(ctx context.Context, user models.Users) error
	Login(ctx context.Context, login models.Login) (string, error)
}
