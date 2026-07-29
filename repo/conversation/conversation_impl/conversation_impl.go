package conversation_impl

import (
	"chatapp-backend/repo/conversation"

	"gorm.io/gorm"
)

var _ conversation.ConversationRepo = (*ConversationImpl)(nil)

type ConversationImpl struct {
	db *gorm.DB
}

func NewConversationImpl(db *gorm.DB) conversation.ConversationRepo {
	return &ConversationImpl{db: db}
}
