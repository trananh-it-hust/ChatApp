package util

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/trananh-it-hust/ChatApp/global"
	"go.uber.org/zap"
)

// Cấu hình upgrader để nâng cấp kết nối HTTP thành WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Cho phép tất cả các nguồn (origin) kết nối
		return true
	},
}

var (
	// Lưu trữ danh sách các client kết nối theo từng conversationID
	clients = make(map[string]map[*websocket.Conn]bool)
	// Mutex để đảm bảo an toàn khi truy cập vào map clients
	mu sync.Mutex
)

// Hàm xử lý kết nối WebSocket
func HandleConnections(c *gin.Context) {
	// Nâng cấp kết nối HTTP thành WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Log.Info("Lỗi WebSocket", zap.Error(err))
		return
	}
	defer ws.Close() // Đóng kết nối WebSocket khi hàm kết thúc

	// Lấy conversationID và senderID từ query parameters
	conversationID := c.Query("conversation_id")
	senderID, err := strconv.Atoi(c.Query("sender_id"))
	if err != nil || conversationID == "" {
		global.Log.Info("Lỗi khi nhận conversationID hoặc senderID", zap.Error(err))
		ws.Close()
		return
	}

	// Ghi log thông tin kết nối WebSocket
	global.Log.Info("WebSocket kết nối thành công", zap.String("ConversationID", conversationID), zap.Int("SenderID", senderID))

	// Thêm kết nối WebSocket vào danh sách clients
	mu.Lock()
	if _, ok := clients[conversationID]; !ok {
		clients[conversationID] = make(map[*websocket.Conn]bool)
	}
	clients[conversationID][ws] = true
	mu.Unlock()

	// Ghi log khi có kết nối WebSocket mới
	global.Log.Info("WebSocket mới tham gia vào conversation", zap.String("ConversationID", conversationID))

	// Đăng ký lắng nghe tin nhắn từ Redis Pub/Sub
	go SubscribeMessages(conversationID, func(msg []byte) {
		mu.Lock()
		for client := range clients[conversationID] {
			// Gửi tin nhắn đến từng client trong conversation
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				// Nếu lỗi, đóng kết nối và xóa client khỏi danh sách
				client.Close()
				delete(clients[conversationID], client)
				global.Log.Info("Lỗi gửi tin nhắn đến client", zap.Error(err))
			}
		}
		mu.Unlock()
	})

	// Lắng nghe tin nhắn từ WebSocket client
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			// Nếu lỗi, xóa client khỏi danh sách và ghi log
			mu.Lock()
			delete(clients[conversationID], ws)
			mu.Unlock()
			global.Log.Info("Lỗi khi đọc tin nhắn từ WebSocket", zap.Error(err))
			break
		}
		// Ghi log tin nhắn nhận được
		global.Log.Info("Tin nhắn nhận được", zap.Int("SenderID", senderID), zap.String("Message", string(msg)))
		// Lưu và phát tin nhắn qua Redis Pub/Sub
		SaveAndPublishMessage(conversationID, int(senderID), string(msg))
	}
}

// Hàm đăng ký lắng nghe tin nhắn từ Redis Pub/Sub
func SubscribeMessages(conversationID string, callback func([]byte)) {
	// Đăng ký vào kênh Redis với tên "chat:<conversationID>"
	pubsub := global.Rdb.Subscribe(context.Background(), fmt.Sprintf("chat:%s", conversationID))
	ch := pubsub.Channel() // Lấy kênh để nhận tin nhắn
	for msg := range ch {
		// Ghi log tin nhắn nhận được từ Redis
		global.Log.Info("Tin nhắn nhận từ Redis", zap.String("ConversationID", conversationID), zap.String("Payload", msg.Payload))
		// Gọi callback để xử lý tin nhắn
		callback([]byte(msg.Payload))
	}
}
