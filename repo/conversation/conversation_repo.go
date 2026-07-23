package conversation

import (
	"chatapp-backend/models"
	"context"
)

type ConversationRepo interface {
	UserExists(ctx context.Context, userID string) (bool, error)
	FindDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error)
	CreateDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error)
	ListDMConversations(ctx context.Context, userID string) ([]models.Conversation, error)
	IsMember(ctx context.Context, conversationID, userID string) (bool, error)
	CreateMessage(ctx context.Context, conversationID, senderID, content string) (*models.Message, error)
	ListMessages(ctx context.Context, conversationID string) ([]models.Message, error)
}
