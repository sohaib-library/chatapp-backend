package conversation_impl

import conversationService "chatapp-backend/service/conversation"

type Handler struct {
	Conversation conversationService.ConversationService
}

func NewHandler(svc conversationService.ConversationService) *Handler {
	return &Handler{Conversation: svc}
}
