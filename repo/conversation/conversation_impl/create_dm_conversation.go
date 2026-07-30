package conversation_impl

import (
	"chatapp-backend/models"
	"chatapp-backend/utils"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (r *ConversationImpl) CreateDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error) {
	var conv models.ConversationDB

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		conv = models.ConversationDB{ID: utils.NewID(), Type: "dm"}
		if err := tx.Create(&conv).Error; err != nil {
			return fmt.Errorf("create conversation: %w", err)
		}

		members := []models.ConversationMember{
			{ConversationID: conv.ID, UserID: userA},
			{ConversationID: conv.ID, UserID: userB},
		}
		if err := tx.Create(&members).Error; err != nil {
			return fmt.Errorf("add conversation members: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Fetch the other user so the caller gets name + id immediately.
	otherUser, err := r.findUser(context.WithoutCancel(ctx), userB)
	if err != nil {
		return nil, err
	}

	return &models.Conversation{
		ID:        conv.ID,
		Type:      conv.Type,
		CreatedAt: conv.CreatedAt,
		OtherUser: otherUser,
	}, nil
}
