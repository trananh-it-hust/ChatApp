package model

import "time"

func (*Conversation) TableName() string {
	return "conversations"
}

type Conversation struct {
	ID        int       `json:"id" db:"id"`
	User1ID   int       `json:"user1_id" db:"user1_id"`
	User2ID   int       `json:"user2_id" db:"user2_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ConversationCreate struct {
	User1ID int `json:"user1_id" db:"user1_id"`
	User2ID int `json:"user2_id" db:"user2_id"`
}

func (c *ConversationCreate) TableName() string {
	return (&Conversation{}).TableName()
}
