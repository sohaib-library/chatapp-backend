package conversation_impl

import (
	"chatapp-backend/models"
	"chatapp-backend/utils"
	"context"
	"fmt"
)

func (r *ConversationImpl) CreateMessage(ctx context.Context, conversationID, senderID, content string) (*models.Message, error) {
	msg := models.MessageDB{
		ID:             utils.NewID(),
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
	}

	if err := r.db.WithContext(ctx).Create(&msg).Error; err != nil {
		return nil, err
	}

	// Fetch sender name so the response is immediately useful.
	var sender models.UserDB
	if err := r.db.WithContext(ctx).
		Select("id, name").
		Where("id = ?", senderID).
		First(&sender).Error; err != nil {
		return nil, fmt.Errorf("fetch sender: %w", err)
	}

	return &models.Message{
		ID:             msg.ID,
		ConversationID: msg.ConversationID,
		SenderID:       msg.SenderID,
		SenderName:     sender.Name,
		Content:        msg.Content,
		CreatedAt:      msg.CreatedAt,
	}, nil
}
