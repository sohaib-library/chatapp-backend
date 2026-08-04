package conversation_impl

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

func (r *ConversationImpl) UsersExist(ctx context.Context, userIDs []string) (bool, error) {
	if len(userIDs) == 0 {
		return false, nil
	}

	var count int64
	result := r.db.WithContext(ctx).
		Model(&models.UserDB{}).
		Where("id IN ?", userIDs).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return int(count) == len(userIDs), nil
}
