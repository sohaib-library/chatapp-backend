package auth_impl

import (
	"chatapp-backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var users models.Users

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid signup request"})
		log.Print(err.Error())
		return
	}

	if err := h.Authuser.SignUp(ctx.Request.Context(), users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": " SignUp Successfully"})
}
