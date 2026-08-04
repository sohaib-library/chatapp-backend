package conversation_impl

import (
	conversationHandler "github.com/sohaib-library/chatapp-backend/handler/conversation"
	conversationService "github.com/sohaib-library/chatapp-backend/service/conversation"
)

var _ conversationHandler.ConversationHandler = (*Handler)(nil)

type Handler struct {
	Conversation conversationService.ConversationService
}

func NewHandler(svc conversationService.ConversationService) *Handler {
	return &Handler{Conversation: svc}
}
