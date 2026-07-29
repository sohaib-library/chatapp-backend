package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
)

func (r *ConversationImpl) ListGroupConversations(ctx context.Context, userID string) ([]models.Conversation, error) {
	var convDBs []models.ConversationDB

	result := r.db.WithContext(ctx).
		Joins("JOIN conversation_members me ON me.conversation_id = conversations.id AND me.user_id = ?", userID).
		Where("conversations.type = ?", "group").
		Order("conversations.created_at DESC").
		Find(&convDBs)

	if result.Error != nil {
		return nil, fmt.Errorf("list group conversations: %w", result.Error)
	}

	conversations := make([]models.Conversation, 0, len(convDBs))
	for _, c := range convDBs {
		members, err := r.listConversationMembers(ctx, c.ID)
		if err != nil {
			return nil, err
		}
		conversations = append(conversations, models.Conversation{
			ID:        c.ID,
			Type:      c.Type,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			Members:   members,
		})
	}

	return conversations, nil
}

func (r *ConversationImpl) listConversationMembers(ctx context.Context, conversationID string) ([]models.UserInfo, error) {
	var rows []models.UserDB

	result := r.db.WithContext(ctx).
		Joins("JOIN conversation_members cm ON cm.user_id = users.id").
		Where("cm.conversation_id = ?", conversationID).
		Order("users.name ASC").
		Find(&rows)

	if result.Error != nil {
		return nil, fmt.Errorf("list conversation members: %w", result.Error)
	}

	members := make([]models.UserInfo, 0, len(rows))
	for _, u := range rows {
		members = append(members, models.UserInfo{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return members, nil
}
