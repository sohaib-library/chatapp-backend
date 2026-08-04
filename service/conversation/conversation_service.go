package conversation

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

type ConversationService interface {
	StartDM(ctx context.Context, currentUserID, otherUserID string) (*models.Conversation, error)
	ListDMs(ctx context.Context, currentUserID string) ([]models.Conversation, error)
	CreateGroup(ctx context.Context, currentUserID, name string, memberIDs []string) (*models.Conversation, error)
	ListGroups(ctx context.Context, currentUserID string) ([]models.Conversation, error)
	SendMessage(ctx context.Context, currentUserID, conversationID, content string) (*models.Message, error)
	ListMessages(ctx context.Context, currentUserID, conversationID string) ([]models.Message, error)
}
