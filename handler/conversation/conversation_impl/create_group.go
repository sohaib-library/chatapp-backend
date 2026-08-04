package conversation_impl

import (
	"net/http"

	"github.com/sohaib-library/chatapp-backend/middleware"
	"github.com/sohaib-library/chatapp-backend/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGroup(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req models.CreateGroupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	group, err := h.Conversation.CreateGroup(ctx.Request.Context(), userID, req.Name, req.MemberIDs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": "Group created",
		"group":   group,
	})
}
