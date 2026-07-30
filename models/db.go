package models

import "time"

// ConversationDB is the GORM model mapped to the conversations table.
type ConversationDB struct {
	ID        string    `gorm:"column:id;primaryKey;type:varchar(36)"`
	Type      string    `gorm:"column:type"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ConversationDB) TableName() string { return "conversations" }

// ConversationMember is the GORM model mapped to the conversation_members table.
type ConversationMember struct {
	ConversationID string `gorm:"column:conversation_id;type:varchar(36)"`
	UserID         string `gorm:"column:user_id"`
}

func (ConversationMember) TableName() string { return "conversation_members" }

// MessageDB is the GORM model mapped to the messages table.
type MessageDB struct {
	ID             string    `gorm:"column:id;primaryKey;type:varchar(36)"`
	ConversationID string    `gorm:"column:conversation_id;type:varchar(36)"`
	SenderID       string    `gorm:"column:sender_id"`
	Content        string    `gorm:"column:content"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

func (MessageDB) TableName() string { return "messages" }

// UserDB is the GORM model mapped to the users table.
type UserDB struct {
	ID    string `gorm:"column:id;primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

func (UserDB) TableName() string { return "users" }
