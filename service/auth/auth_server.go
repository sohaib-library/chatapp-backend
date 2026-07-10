package auth

import (
	
	"chatapp-backend/models"
	"context"
)


type AuthService interface {
	SignUp(ctx context.Context, user models.Users) error
}
