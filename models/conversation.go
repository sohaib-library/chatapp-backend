package models

import "time"

// Conversation is the API response model for a conversation.
type Conversation struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	Name      string     `json:"name,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	OtherUser *UserInfo  `json:"other_user,omitempty"`
	Members   []UserInfo `json:"members,omitempty"`
}

// UserInfo is the API response model for a user inside a conversation.
type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Message is the API response model for a message.
type Message struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	SenderID       string    `json:"sender_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

// CreateDMRequest is the request body for starting a DM.
type CreateDMRequest struct {
	UserID string `json:"user_id"`
}

// CreateGroupRequest is the request body for creating a group.
type CreateGroupRequest struct {
	Name      string   `json:"name"`
	MemberIDs []string `json:"member_ids"`
}

// SendMessageRequest is the request body for sending a message.
type SendMessageRequest struct {
	Content string `json:"content"`
}
