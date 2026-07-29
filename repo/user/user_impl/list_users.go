package user_impl

import (
	"chatapp-backend/models"
	"context"
)

func (r *UserImpl) ListUsers(ctx context.Context, excludeUserID string) ([]models.UserInfo, error) {
	var rows []models.UserDB

	result := r.db.WithContext(ctx).
		Where("id <> ?", excludeUserID).
		Order("name ASC").
		Find(&rows)

	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]models.UserInfo, 0, len(rows))
	for _, u := range rows {
		users = append(users, models.UserInfo{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return users, nil
}
