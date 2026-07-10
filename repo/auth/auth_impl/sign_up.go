package auth_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"context"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

func ( r *AuthImpl ) CreateUser (ctx context.Context, user models.Users) error {


	_, err := r.db.ExecContext (ctx, `

		INSERT INTO users (
		name,
		email,
		password )
		VALUES 
		($1,$2,$3)
	`, 
	user.Name,
	user.Email, 
	user.Password)


	if err != nil {

		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return apperror.ErrUserExists
		}

		return fmt.Errorf("insert user: %w", err)

	}

	return nil
}


func (r *AuthImpl) UserExists(ctx context.Context, email string) (bool, error) {

	var exists bool

	err := r.db.QueryRowContext (ctx, `

		SELECT EXISTS(
			SELECT 1 FROM users 
			WHERE
			email = $1
		)
	`,
	 email).

	 Scan(&exists)
	 
	if err != nil {
		return false, fmt.Errorf("check user exists: %w", err)
	}


	return exists, nil
}
