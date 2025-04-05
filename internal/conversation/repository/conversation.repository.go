package repository

import (
	"github.com/trananh-it-hust/ChatApp/global"
	"github.com/trananh-it-hust/ChatApp/internal/conversation/model"
	"go.uber.org/zap"
)

type ConversationRepository interface {
	CreateConversation(user1ID, user2ID int) (int, error)
	GetConversationByID(id int) (model.Conversation, error)
	GetConversationsByUserID(userID int) ([]model.Conversation, error)
	GetConversationByUserIDs(user1ID, user2ID int) (model.Conversation, error)
}

type ConversationRepositoryImpl struct{}

func NewConversationRepository() ConversationRepository {
	return &ConversationRepositoryImpl{}
}

func (cr *ConversationRepositoryImpl) CreateConversation(user1ID, user2ID int) (int, error) {
	db := global.MDB
	conversation := model.Conversation{
		User1ID: user1ID,
		User2ID: user2ID,
	}
	if err := db.Create(&conversation).Error; err != nil {
		global.Log.Error("Failed to create conversation:", zap.Error(err))
		return 0, err
	}
	return conversation.ID, nil
}

func (cr *ConversationRepositoryImpl) GetConversationByID(id int) (model.Conversation, error) {
	db := global.MDB
	var conversation model.Conversation
	if err := db.Where("id = ?", id).First(&conversation).Error; err != nil {
		global.Log.Error("Failed to get conversation by ID:", zap.Error(err))
		return model.Conversation{}, err
	}
	return conversation, nil
}
func (cr *ConversationRepositoryImpl) GetConversationsByUserID(userID int) ([]model.Conversation, error) {
	db := global.MDB
	var conversations []model.Conversation
	if err := db.Where("user1_id = ? OR user2_id = ?", userID, userID).Find(&conversations).Error; err != nil {
		global.Log.Error("Failed to get conversations by user ID:", zap.Error(err))
		return nil, err
	}
	return conversations, nil
}

func (cr *ConversationRepositoryImpl) GetConversationByUserIDs(user1ID, user2ID int) (model.Conversation, error) {
	db := global.MDB
	var conversation model.Conversation
	if err := db.Where("(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)", user1ID, user2ID, user2ID, user1ID).First(&conversation).Error; err != nil {
		global.Log.Error("Failed to get conversation by user IDs:", zap.Error(err))
		return model.Conversation{}, err
	}
	return conversation, nil
}
