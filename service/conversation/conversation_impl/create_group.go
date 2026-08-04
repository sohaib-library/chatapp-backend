package conversation_impl

import (
	"context"
	"strings"

	apperror "github.com/sohaib-library/chatapp-backend/error"
	"github.com/sohaib-library/chatapp-backend/models"
)

func (s *ConversationServiceImpl) CreateGroup(ctx context.Context, currentUserID, name string, memberIDs []string) (*models.Conversation, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, apperror.ErrInvalidGroupName
	}

	uniqueMembers := make(map[string]struct{})
	cleanedMembers := make([]string, 0, len(memberIDs)+1)

	uniqueMembers[currentUserID] = struct{}{}
	cleanedMembers = append(cleanedMembers, currentUserID)

	for _, memberID := range memberIDs {
		memberID = strings.TrimSpace(memberID)
		if memberID == "" || memberID == currentUserID {
			continue
		}
		if _, exists := uniqueMembers[memberID]; exists {
			continue
		}
		uniqueMembers[memberID] = struct{}{}
		cleanedMembers = append(cleanedMembers, memberID)
	}

	if len(cleanedMembers) < 2 {
		return nil, apperror.ErrInvalidGroupMembers
	}

	exists, err := s.repo.UsersExist(ctx, cleanedMembers)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, apperror.ErrUserNotFound
	}

	return s.repo.CreateGroupConversation(ctx, name, cleanedMembers)
}
