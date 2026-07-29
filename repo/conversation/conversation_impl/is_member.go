package conversation_impl

import (
	"chatapp-backend/models"
	"context"
)

func (r *ConversationImpl) IsMember(ctx context.Context, conversationID, userID string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).
		Model(&models.ConversationMember{}).
		Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
