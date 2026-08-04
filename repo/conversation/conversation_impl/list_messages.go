package conversation_impl

import (
	"context"
	"fmt"

	"github.com/sohaib-library/chatapp-backend/models"
)

// msgWithSender is a flat scan target for the JOIN query.
type msgWithSender struct {
	ID             string `gorm:"column:id"`
	ConversationID string `gorm:"column:conversation_id"`
	SenderID       string `gorm:"column:sender_id"`
	SenderName     string `gorm:"column:sender_name"`
	Content        string `gorm:"column:content"`
	CreatedAt      string `gorm:"column:created_at"`
}

func (r *ConversationImpl) ListMessages(ctx context.Context, conversationID string) ([]models.Message, error) {
	var rows []msgWithSender

	result := r.db.WithContext(ctx).Raw(`
		SELECT
			m.id,
			m.conversation_id,
			m.sender_id,
			u.name  AS sender_name,
			m.content,
			m.created_at
		FROM messages m
		JOIN users u ON u.id = m.sender_id
		WHERE m.conversation_id = ?
		ORDER BY m.created_at ASC
	`, conversationID).Scan(&rows)

	if result.Error != nil {
		return nil, fmt.Errorf("list messages: %w", result.Error)
	}

	messages := make([]models.Message, 0, len(rows))
	for _, row := range rows {
		messages = append(messages, models.Message{
			ID:             row.ID,
			ConversationID: row.ConversationID,
			SenderID:       row.SenderID,
			SenderName:     row.SenderName,
			Content:        row.Content,
		})
	}

	return messages, nil
}
