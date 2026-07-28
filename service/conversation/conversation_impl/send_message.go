package conversation_impl

import (
	apperror "chatapp-backend/error"
	"chatapp-backend/models"
	"context"
	"strings"
)

func (s *ConversationServiceImpl) SendMessage(ctx context.Context, currentUserID, conversationID, content string) (*models.Message, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, apperror.ErrInvalidMessage
	}

	isMember, err := s.repo.IsMember(ctx, conversationID, currentUserID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, apperror.ErrNotConversationMember
	}

	message, err := s.repo.CreateMessage(ctx, conversationID, currentUserID, content)
	if err != nil {
		return nil, err
	}

	if s.notifier != nil {
		memberIDs, memberErr := s.repo.ListMemberIDs(ctx, conversationID)
		if memberErr == nil {
			s.notifier.NotifyMessage(conversationID, memberIDs, message)
		}
	}

	return message, nil
}
