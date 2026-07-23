package conversation_impl

import (
	"chatapp-backend/models"
	"context"
)

func (s *ConversationServiceImpl) ListDMs(ctx context.Context, currentUserID string) ([]models.Conversation, error) {
	return s.repo.ListDMConversations(ctx, currentUserID)
}
