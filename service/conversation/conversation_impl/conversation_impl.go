package conversation_impl

import (
	"github.com/sohaib-library/chatapp-backend/repo/conversation"
	conversationService "github.com/sohaib-library/chatapp-backend/service/conversation"
)

var _ conversationService.ConversationService = (*ConversationServiceImpl)(nil)

type ConversationServiceImpl struct {
	repo     conversation.ConversationRepo
	notifier conversationService.RealtimeNotifier
}

func NewConversation(
	repo conversation.ConversationRepo,
	notifier conversationService.RealtimeNotifier,
) conversationService.ConversationService {
	return &ConversationServiceImpl{
		repo:     repo,
		notifier: notifier,
	}
}
