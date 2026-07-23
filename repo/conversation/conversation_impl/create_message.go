package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) CreateMessage(ctx context.Context, conversationID, senderID, content string) (*models.Message, error) {
	var msg models.Message
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO messages (conversation_id, sender_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, conversation_id, sender_id, content, created_at
	`, conversationID, senderID, content).Scan(
		&msg.ID,
		&msg.ConversationID,
		&msg.SenderID,
		&msg.Content,
		&msg.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}
	return &msg, nil
}
