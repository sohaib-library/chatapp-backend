package conversation_impl

import (
	"net/http"

	"github.com/sohaib-library/chatapp-backend/middleware"
	"github.com/sohaib-library/chatapp-backend/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StartDM(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req models.CreateDMRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	conversation, err := h.Conversation.StartDM(ctx.Request.Context(), userID, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success":      "DM conversation ready",
		"conversation": conversation,
	})
}
