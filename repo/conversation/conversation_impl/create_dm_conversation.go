package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) CreateDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin dm conversation tx: %w", err)
	}
	defer tx.Rollback()

	var conv models.Conversation
	err = tx.QueryRowContext(ctx, `
		INSERT INTO conversations (type)
		VALUES ('dm')
		RETURNING id, type, created_at
	`).Scan(&conv.ID, &conv.Type, &conv.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("create conversation: %w", err)
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO conversation_members (conversation_id, user_id)
		VALUES ($1, $2), ($1, $3)
	`, conv.ID, userA, userB)
	if err != nil {
		return nil, fmt.Errorf("add conversation members: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit dm conversation: %w", err)
	}

	return &conv, nil
}
