package conversation_impl

import (
	"net/http"

	"chatapp-backend/middleware"
	"chatapp-backend/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendMessage(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	conversationID := ctx.Param("id")
	if conversationID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "conversation id is required"})
		return
	}

	var req models.SendMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	message, err := h.Conversation.SendMessage(ctx.Request.Context(), userID, conversationID, req.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": "Message sent",
		"message": message,
	})
}
