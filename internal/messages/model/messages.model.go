package model

import "time"

type Message struct {
	ID             int64     `json:"id"`
	ConversationID int64     `json:"conversation_id"`
	SenderID       int64     `json:"sender_id"`
	Content        string    `json:"content"`
	IsRead         bool      `json:"is_read"`
	CreatedAt      time.Time `json:"created_at"`
}

func (m *Message) TableName() string {
	return "messages"
}

type MessageCreate struct {
	ConversationID int64  `json:"conversation_id"`
	SenderID       int64  `json:"sender_id"`
	Content        string `json:"content"`
}

func (m *MessageCreate) TableName() string {
	return (&Message{}).TableName()
}
