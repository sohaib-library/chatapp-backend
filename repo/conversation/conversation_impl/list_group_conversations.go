package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) ListGroupConversations(ctx context.Context, userID string) ([]models.Conversation, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT c.id, c.type, c.name, c.created_at
		FROM conversations c
		JOIN conversation_members me ON me.conversation_id = c.id AND me.user_id = $1
		WHERE c.type = 'group'
		ORDER BY c.created_at DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("list group conversations: %w", err)
	}
	defer rows.Close()

	conversations := make([]models.Conversation, 0)
	for rows.Next() {
		var conv models.Conversation
		if err := rows.Scan(&conv.ID, &conv.Type, &conv.Name, &conv.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan group conversation: %w", err)
		}

		members, err := r.listConversationMembers(ctx, conv.ID)
		if err != nil {
			return nil, err
		}
		conv.Members = members
		conversations = append(conversations, conv)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate group conversations: %w", err)
	}

	return conversations, nil
}

func (r *ConversationImpl) listConversationMembers(ctx context.Context, conversationID string) ([]models.UserInfo, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT u.id, u.name, u.email
		FROM conversation_members cm
		JOIN users u ON u.id = cm.user_id
		WHERE cm.conversation_id = $1
		ORDER BY u.name ASC
	`, conversationID)
	if err != nil {
		return nil, fmt.Errorf("list conversation members: %w", err)
	}
	defer rows.Close()

	members := make([]models.UserInfo, 0)
	for rows.Next() {
		var member models.UserInfo
		if err := rows.Scan(&member.ID, &member.Name, &member.Email); err != nil {
			return nil, fmt.Errorf("scan conversation member: %w", err)
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate conversation members: %w", err)
	}

	return members, nil
}
