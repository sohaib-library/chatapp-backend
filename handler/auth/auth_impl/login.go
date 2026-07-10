package auth_impl

import (
	"chatapp-backend/models"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {

	var login models.Login


	var emailRegex = regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		log.Print(err.Error())
		return
	}


	if !emailRegex.MatchString(login.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{

			"error": "Enter Valid Email",
		})

		return
	}


}