package conversation_impl

import (
	"chatapp-backend/models"
	"context"
)

func (s *ConversationServiceImpl) ListGroups(ctx context.Context, currentUserID string) ([]models.Conversation, error) {
	return s.repo.ListGroupConversations(ctx, currentUserID)
}
