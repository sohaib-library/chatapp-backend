package conversation_impl

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

func (r *ConversationImpl) UserExists(ctx context.Context, userID string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).
		Model(&models.UserDB{}).
		Where("id = ?", userID).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
