package util

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"main.go/global"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients = make(map[string]map[*websocket.Conn]bool)
	mu      sync.Mutex
)

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Log.Info("Lỗi WebSocket", zap.Error(err))
		return
	}
	defer ws.Close()

	conversationID := c.Query("conversation_id")
	senderID, err := strconv.Atoi(c.Query("sender_id"))
	if err != nil || conversationID == "" {
		global.Log.Info("Lỗi khi nhận conversationID hoặc senderID", zap.Error(err))
		ws.Close()
		return
	}

	// Log thông tin kết nối WebSocket
	global.Log.Info("WebSocket kết nối thành công", zap.String("ConversationID", conversationID), zap.Int("SenderID", senderID))

	mu.Lock()
	if _, ok := clients[conversationID]; !ok {
		clients[conversationID] = make(map[*websocket.Conn]bool)
	}
	clients[conversationID][ws] = true
	mu.Unlock()

	// Log khi kết nối WebSocket mới được thêm vào
	global.Log.Info("WebSocket mới tham gia vào conversation", zap.String("ConversationID", conversationID))

	// Subscribe Redis Pub/Sub
	go SubscribeMessages(conversationID, func(msg []byte) {
		mu.Lock()
		for client := range clients[conversationID] {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(clients[conversationID], client)
				global.Log.Info("Lỗi gửi tin nhắn đến client", zap.Error(err))
			}
		}
		mu.Unlock()
	})

	// Nhận tin nhắn từ WebSocket client
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients[conversationID], ws)
			mu.Unlock()
			global.Log.Info("Lỗi khi đọc tin nhắn từ WebSocket", zap.Error(err))
			break
		}
		global.Log.Info("Tin nhắn nhận được", zap.Int("SenderID", senderID), zap.String("Message", string(msg)))
		SaveAndPublishMessage(conversationID, int(senderID), string(msg))
	}
}

func SubscribeMessages(conversationID string, callback func([]byte)) {
	pubsub := global.Rdb.Subscribe(context.Background(), fmt.Sprintf("chat:%s", conversationID))
	ch := pubsub.Channel()
	for msg := range ch {
		global.Log.Info("Tin nhắn nhận từ Redis", zap.String("ConversationID", conversationID), zap.String("Payload", msg.Payload))
		callback([]byte(msg.Payload))
	}
}
