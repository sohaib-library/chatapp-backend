package conversation_impl

import (
	"net/http"

	"github.com/sohaib-library/chatapp-backend/middleware"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListMessages(ctx *gin.Context) {
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

	messages, err := h.Conversation.ListMessages(ctx.Request.Context(), userID, conversationID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
