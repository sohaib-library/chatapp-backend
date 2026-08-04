package conversation_impl

import (
	"context"
	"strings"

	apperror "github.com/sohaib-library/chatapp-backend/error"
	"github.com/sohaib-library/chatapp-backend/models"
)

func (s *ConversationServiceImpl) StartDM(ctx context.Context, currentUserID, otherUserID string) (*models.Conversation, error) {
	otherUserID = strings.TrimSpace(otherUserID)
	if otherUserID == "" {
		return nil, apperror.ErrInvalidRequest
	}

	if currentUserID == otherUserID {
		return nil, apperror.ErrCannotDMYourself
	}

	exists, err := s.repo.UserExists(ctx, otherUserID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, apperror.ErrUserNotFound
	}

	existing, err := s.repo.FindDMConversation(ctx, currentUserID, otherUserID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}

	return s.repo.CreateDMConversation(ctx, currentUserID, otherUserID)
}
