package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service MessageService
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		service: NewMessageService(),
	}
}

func (h *MessageHandler) GetMessage(c *gin.Context) {
	message := h.service.GetService()
	c.IndentedJSON(http.StatusOK, message)
}
