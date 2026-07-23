package conversation_impl

import (
	"net/http"

	"chatapp-backend/middleware"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListDMs(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	conversations, err := h.Conversation.ListDMs(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"conversations": conversations})
}
