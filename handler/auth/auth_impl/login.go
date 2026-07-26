package auth_impl

import (
	"chatapp-backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	var login models.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		log.Print(err.Error())
		return
	}

	token, err := h.Authuser.Login(ctx.Request.Context(), login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": "Login Successfully",
		"Token":   token,
	})
}
