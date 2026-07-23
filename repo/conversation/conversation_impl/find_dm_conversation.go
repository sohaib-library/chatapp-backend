package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *ConversationImpl) FindDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error) {
	var conv models.Conversation
	var name sql.NullString

	err := r.db.QueryRowContext(ctx, `
		SELECT c.id, c.type, c.name, c.created_at
		FROM conversations c
		JOIN conversation_members m1 ON m1.conversation_id = c.id AND m1.user_id = $1
		JOIN conversation_members m2 ON m2.conversation_id = c.id AND m2.user_id = $2
		WHERE c.type = 'dm'
		LIMIT 1
	`, userA, userB).Scan(&conv.ID, &conv.Type, &name, &conv.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("find dm conversation: %w", err)
	}

	if name.Valid {
		conv.Name = name.String
	}

	return &conv, nil
}
