package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) ListMessages(ctx context.Context, conversationID string) ([]models.Message, error) {
	var msgDBs []models.MessageDB

	result := r.db.WithContext(ctx).
		Where("conversation_id = ?", conversationID).
		Order("created_at ASC").
		Find(&msgDBs)

	if result.Error != nil {
		return nil, fmt.Errorf("list messages: %w", result.Error)
	}

	messages := make([]models.Message, 0, len(msgDBs))
	for _, m := range msgDBs {
		messages = append(messages, models.Message{
			ID:             m.ID,
			ConversationID: m.ConversationID,
			SenderID:       m.SenderID,
			Content:        m.Content,
			CreatedAt:      m.CreatedAt,
		})
	}

	return messages, nil
}
