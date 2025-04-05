package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	converC "github.com/trananh-it-hust/ChatApp/internal/conversation/controller"
	userC "github.com/trananh-it-hust/ChatApp/internal/user/controller"
	"github.com/trananh-it-hust/ChatApp/pkg/util"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong"})
	})

	user := r.Group("/auth")
	{
		user.POST("/register", (userC.NewUserController()).CreateUser)
		user.POST("/login", (userC.NewUserController()).LoginUser)
	}
	conver := r.Group("/conversation")
	{
		conver.POST("/create", (converC.NewConversationController()).CreateConversation)
	}
	// WebSocket route
	r.GET("/ws", func(c *gin.Context) {
		util.HandleConnections(c)
	})
	return r
}
