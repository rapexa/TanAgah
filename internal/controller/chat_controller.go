package controller

import (
	"TanAgah/internal/model"
	"TanAgah/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

type MessageController struct {
	MessageService service.MessageService
}

func NewMessageController(MessageService service.MessageService) *MessageController {
	return &MessageController{MessageService}
}

type MessagePayload struct {
	Type      string `json:"type"`
	Sender    uint   `json:"sender"`
	Receiver  uint   `json:"receiver"`
	Content   string `json:"content"`
	MessageID uint   `json:"message_id,omitempty"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]uint)
var mu sync.Mutex

func (cm *MessageController) GetChatHistory(c *gin.Context) {
	senderID, _ := strconv.ParseUint(c.Param("sender_id"), 10, 32)
	receiverID, _ := strconv.ParseUint(c.Param("receiver_id"), 10, 32)

	messages, err := cm.MessageService.GetChatHistory(uint(senderID), uint(receiverID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chat history"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (cm *MessageController) ChatWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WebSocket upgrade failed"})
		return
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			break
		}

		var payload MessagePayload
		json.Unmarshal(msg, &payload)

		switch payload.Type {
		case "message":
			cm.MessageService.SaveMessageDb(&model.Message{
				SenderID:   payload.Sender,
				ReceiverID: payload.Receiver,
				Content:    payload.Content,
			})
		case "edit":
			cm.MessageService.EditMessageDb(payload.MessageID, payload.Content, payload.Sender)
		case "delete":
			cm.MessageService.DeleteMessageDb(payload.MessageID, payload.Sender)
		}

		mu.Lock()
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, msg)
		}
		mu.Unlock()
	}
}
