package conversation

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

type ConversationRepo interface {
	UserExists(ctx context.Context, userID string) (bool, error)
	UsersExist(ctx context.Context, userIDs []string) (bool, error)
	FindDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error)
	CreateDMConversation(ctx context.Context, userA, userB string) (*models.Conversation, error)
	ListDMConversations(ctx context.Context, userID string) ([]models.Conversation, error)
	CreateGroupConversation(ctx context.Context, name string, memberIDs []string) (*models.Conversation, error)
	ListGroupConversations(ctx context.Context, userID string) ([]models.Conversation, error)
	IsMember(ctx context.Context, conversationID, userID string) (bool, error)
	ListMemberIDs(ctx context.Context, conversationID string) ([]string, error)
	CreateMessage(ctx context.Context, conversationID, senderID, content string) (*models.Message, error)
	ListMessages(ctx context.Context, conversationID string) ([]models.Message, error)
}
