package user_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *UserImpl) ListUsers(ctx context.Context, excludeUserID string) ([]models.UserInfo, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, email
		FROM users
		WHERE id <> $1
		ORDER BY name ASC
	`, excludeUserID)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	defer rows.Close()

	users := make([]models.UserInfo, 0)
	for rows.Next() {
		var user models.UserInfo
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate users: %w", err)
	}

	return users, nil
}
