package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (r *ConversationImpl) CreateGroupConversation(ctx context.Context, name string, memberIDs []string) (*models.Conversation, error) {
	var conv models.ConversationDB

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		conv = models.ConversationDB{Type: "group", Name: name}
		if err := tx.Create(&conv).Error; err != nil {
			return fmt.Errorf("create group conversation: %w", err)
		}

		members := make([]models.ConversationMember, len(memberIDs))
		for i, id := range memberIDs {
			members[i] = models.ConversationMember{ConversationID: conv.ID, UserID: id}
		}
		if err := tx.Create(&members).Error; err != nil {
			return fmt.Errorf("add group members: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	memberInfos, err := r.listConversationMembers(context.WithoutCancel(ctx), conv.ID)
	if err != nil {
		return nil, err
	}

	return &models.Conversation{
		ID:        conv.ID,
		Type:      conv.Type,
		Name:      conv.Name,
		CreatedAt: conv.CreatedAt,
		Members:   memberInfos,
	}, nil
}
