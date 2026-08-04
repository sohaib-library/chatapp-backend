package conversation_impl

import (
	"context"
	"errors"

	"github.com/sohaib-library/chatapp-backend/models"

	"gorm.io/gorm"
)

func (r *ConversationImpl) FindDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error) {
	var conv models.ConversationDB

	result := r.db.WithContext(ctx).
		Joins("JOIN conversation_members m1 ON m1.conversation_id = conversations.id AND m1.user_id = ?", userA).
		Joins("JOIN conversation_members m2 ON m2.conversation_id = conversations.id AND m2.user_id = ?", userB).
		Where("conversations.type = ?", "dm").
		First(&conv)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	// Fetch the other user (userB from userA's perspective)
	otherUser, err := r.findUser(ctx, userB)
	if err != nil {
		return nil, err
	}

	return &models.Conversation{
		ID:        conv.ID,
		Type:      conv.Type,
		Name:      conv.Name,
		CreatedAt: conv.CreatedAt,
		OtherUser: otherUser,
	}, nil
}

// findUser fetches a single user's info by ID.
func (r *ConversationImpl) findUser(ctx context.Context, userID string) (*models.UserInfo, error) {
	var u models.UserDB
	if err := r.db.WithContext(ctx).
		Select("id, name, email").
		Where("id = ?", userID).
		First(&u).Error; err != nil {
		return nil, err
	}
	return &models.UserInfo{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
