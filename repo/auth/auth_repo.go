package auth

import "context"

type AuthRepo interface {
	CreateUser(ctx context.Context, name, email, password string) error
	UserExists(ctx context.Context, name, email string) (bool, error)
}
