package conversation_impl

import (
	"chatapp-backend/repo/conversation"
	conversationService "chatapp-backend/service/conversation"
)

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
