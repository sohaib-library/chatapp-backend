package conversation_impl

import (
	"chatapp-backend/models"
	"context"
)

func (r *ConversationImpl) CreateMessage(ctx context.Context, conversationID, senderID, content string) (*models.Message, error) {
	msg := models.MessageDB{
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
	}

	result := r.db.WithContext(ctx).Create(&msg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.Message{
		ID:             msg.ID,
		ConversationID: msg.ConversationID,
		SenderID:       msg.SenderID,
		Content:        msg.Content,
		CreatedAt:      msg.CreatedAt,
	}, nil
}
