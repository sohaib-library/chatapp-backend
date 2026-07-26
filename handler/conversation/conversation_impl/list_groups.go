package conversation_impl

import (
	"net/http"

	"chatapp-backend/middleware"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListGroups(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	groups, err := h.Conversation.ListGroups(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}
