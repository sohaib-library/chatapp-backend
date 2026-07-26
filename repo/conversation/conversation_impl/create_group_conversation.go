package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) CreateGroupConversation(ctx context.Context, name string, memberIDs []string) (*models.Conversation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin group conversation tx: %w", err)
	}
	defer tx.Rollback()

	var conv models.Conversation
	err = tx.QueryRowContext(ctx, `
		INSERT INTO conversations (type, name)
		VALUES ('group', $1)
		RETURNING id, type, name, created_at
	`, name).Scan(&conv.ID, &conv.Type, &conv.Name, &conv.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("create group conversation: %w", err)
	}

	for _, memberID := range memberIDs {
		if _, err := tx.ExecContext(ctx, `
			INSERT INTO conversation_members (conversation_id, user_id)
			VALUES ($1, $2)
		`, conv.ID, memberID); err != nil {
			return nil, fmt.Errorf("add group member: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit group conversation: %w", err)
	}

	members, err := r.listConversationMembers(ctx, conv.ID)
	if err != nil {
		return nil, err
	}
	conv.Members = members

	return &conv, nil
}
