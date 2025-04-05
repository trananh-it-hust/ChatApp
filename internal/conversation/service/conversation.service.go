package service

import "github.com/trananh-it-hust/ChatApp/internal/conversation/repository"

type ConversationService interface {
	CreateConversation(user1ID, user2ID int) (int, error)
}

type ConversationServiceImpl struct {
	ConversationRepository repository.ConversationRepository
}

func NewConversationService() ConversationService {
	return &ConversationServiceImpl{
		ConversationRepository: repository.NewConversationRepository(),
	}
}

func (cs *ConversationServiceImpl) CreateConversation(user1ID, user2ID int) (int, error) {
	conversationID, err := cs.ConversationRepository.CreateConversation(user1ID, user2ID)
	if err != nil {
		return 0, err
	}
	return conversationID, nil
}
