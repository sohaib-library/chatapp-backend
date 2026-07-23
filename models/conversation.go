package models

import "time"

type Conversation struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	OtherUser *UserInfo `json:"other_user,omitempty"`
}

type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Message struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateDMRequest struct {
	UserID string `json:"user_id"`
}

type SendMessageRequest struct {
	Content string `json:"content"`
}
