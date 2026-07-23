package conversation_impl

import (
	"chatapp-backend/repo/conversation"
	conversationService "chatapp-backend/service/conversation"
)

type ConversationServiceImpl struct {
	repo conversation.ConversationRepo
}

func NewConversation(repo conversation.ConversationRepo) conversationService.ConversationService {
	return &ConversationServiceImpl{repo: repo}
}
