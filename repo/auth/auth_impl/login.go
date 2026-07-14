package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *AuthImpl) Login(ctx context.Context, email string) (*models.Users, error) {

	var user models.Users

	err := r.db.QueryRowContext(ctx, `

		SELECT id,
		name,
		email,
		password
		FROM users
		WHERE email = $1
	`, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	return &user, nil
}
