package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trananh-it-hust/ChatApp/internal/conversation/model"
	"github.com/trananh-it-hust/ChatApp/internal/conversation/service"
	"github.com/trananh-it-hust/ChatApp/pkg/response"
)

type ConversationController struct {
	ConversationService service.ConversationService
}

func NewConversationController() *ConversationController {
	return &ConversationController{
		ConversationService: service.NewConversationService(),
	}
}
func (cc *ConversationController) CreateConversation(ctx *gin.Context) {
	var request model.ConversationCreate

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Invalid request")))
		return
	}
	if request.User1ID == 0 || request.User2ID == 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Missing required fields")))
		return
	}
	if request.User1ID == request.User2ID {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors.New("Cannot create conversation with the same user")))
		return
	}
	user1ID := request.User1ID
	user2ID := request.User2ID

	conversationID, err := cc.ConversationService.CreateConversation(user1ID, user2ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Conversation created successfully", gin.H{"conversation_id": conversationID}))
}
