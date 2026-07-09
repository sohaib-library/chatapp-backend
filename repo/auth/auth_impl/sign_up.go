package auth_impl

import (
	"context"
	"fmt"
)

func (r *AuthImpl) CreateUser(ctx context.Context, name, email, password string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`, name, email, password)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

func (r *AuthImpl) UserExists(ctx context.Context, name, email string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE name = $1 OR email = $2
		)
	`, name, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check user exists: %w", err)
	}

	return exists, nil
}
