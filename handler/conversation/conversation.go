package conversation

import "github.com/gin-gonic/gin"

type ConversationHandler interface {
	StartDM(ctx *gin.Context)
	ListDMs(ctx *gin.Context)
	CreateGroup(ctx *gin.Context)
	ListGroups(ctx *gin.Context)
	SendMessage(ctx *gin.Context)
	ListMessages(ctx *gin.Context)
}
