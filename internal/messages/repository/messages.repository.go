package repository

import (
	"strconv"

	"github.com/trananh-it-hust/ChatApp/global"
	"github.com/trananh-it-hust/ChatApp/internal/messages/model"
)

type MessageRepository interface {
	CreateMessage(userID int, conversationID, content string) (int, error)
}

type MessageRepositoryImpl struct{}

func NewMessageRepository() MessageRepository {
	return &MessageRepositoryImpl{}
}

func (mr *MessageRepositoryImpl) CreateMessage(userID int, conversationID, content string) (int, error) {
	db := global.MDB
	conversationIDInt, err := strconv.ParseInt(conversationID, 10, 64)
	if err != nil {
		return 0, err
	}
	message := model.Message{
		ConversationID: int64(conversationIDInt),
		SenderID:       int64(userID),
		Content:        content,
	}
	if err := db.Create(&message).Error; err != nil {
		return 0, err
	}
	return int(message.ID), nil
}
