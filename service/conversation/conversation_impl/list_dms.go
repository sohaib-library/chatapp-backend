package conversation_impl

import (
	"context"

	"github.com/sohaib-library/chatapp-backend/models"
)

func (s *ConversationServiceImpl) ListDMs(ctx context.Context, currentUserID string) ([]models.Conversation, error) {
	return s.repo.ListDMConversations(ctx, currentUserID)
}
