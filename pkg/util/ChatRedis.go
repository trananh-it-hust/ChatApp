package util

import (
	"context"
	"fmt"

	"github.com/trananh-it-hust/ChatApp/global"
	"github.com/trananh-it-hust/ChatApp/internal/messages/repository"
)

func SaveAndPublishMessage(conversationID string, senderID int, message string) error {

	messRepo := repository.NewMessageRepository()
	if _, error := messRepo.CreateMessage(senderID, conversationID, message); error != nil {
		return fmt.Errorf("failed to save message: %v", error)
	}

	channel := fmt.Sprintf("chat:%s", conversationID)
	messageData := fmt.Sprintf("%d:%s", senderID, message)
	return global.Rdb.Publish(context.Background(), channel, messageData).Err()
}
