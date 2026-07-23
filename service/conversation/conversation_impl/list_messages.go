package conversation_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"context"
)

func (s *ConversationServiceImpl) ListMessages(ctx context.Context, currentUserID, conversationID string) ([]models.Message, error) {
	isMember, err := s.repo.IsMember(ctx, conversationID, currentUserID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, apperror.ErrNotConversationMember
	}

	return s.repo.ListMessages(ctx, conversationID)
}
