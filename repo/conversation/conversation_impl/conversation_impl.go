package conversation_impl

import (
	"database/sql"

	"chatapp-backend/repo/conversation"
)

type ConversationImpl struct {
	db *sql.DB
}

func NewConversationImpl(db *sql.DB) conversation.ConversationRepo {
	return &ConversationImpl{db: db}
}
