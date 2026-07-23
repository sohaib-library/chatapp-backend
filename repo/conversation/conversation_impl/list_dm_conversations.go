package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) ListDMConversations(ctx context.Context, userID string) ([]models.Conversation, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			c.id,
			c.type,
			c.created_at,
			u.id,
			u.name,
			u.email
		FROM conversations c
		JOIN conversation_members me ON me.conversation_id = c.id AND me.user_id = $1
		JOIN conversation_members other ON other.conversation_id = c.id AND other.user_id <> $1
		JOIN users u ON u.id = other.user_id
		WHERE c.type = 'dm'
		ORDER BY c.created_at DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("list dm conversations: %w", err)
	}
	defer rows.Close()

	conversations := make([]models.Conversation, 0)
	for rows.Next() {
		var conv models.Conversation
		var other models.UserInfo
		if err := rows.Scan(
			&conv.ID,
			&conv.Type,
			&conv.CreatedAt,
			&other.ID,
			&other.Name,
			&other.Email,
		); err != nil {
			return nil, fmt.Errorf("scan dm conversation: %w", err)
		}
		conv.OtherUser = &other
		conversations = append(conversations, conv)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate dm conversations: %w", err)
	}

	return conversations, nil
}
