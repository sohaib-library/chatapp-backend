package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)

}