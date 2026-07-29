package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *ConversationImpl) FindDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error) {
	var conv models.ConversationDB

	result := r.db.WithContext(ctx).
		Joins("JOIN conversation_members m1 ON m1.conversation_id = conversations.id AND m1.user_id = ?", userA).
		Joins("JOIN conversation_members m2 ON m2.conversation_id = conversations.id AND m2.user_id = ?", userB).
		Where("conversations.type = ?", "dm").
		First(&conv)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &models.Conversation{
		ID:        conv.ID,
		Type:      conv.Type,
		Name:      conv.Name,
		CreatedAt: conv.CreatedAt,
	}, nil
}
