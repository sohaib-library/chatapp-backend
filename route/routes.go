package route

import (
	"database/sql"

	auth_handler "chatapp-backend/handler/auth/auth_impl"
	conversation_handler "chatapp-backend/handler/conversation/conversation_impl"
	"chatapp-backend/middleware"
	auth_repo "chatapp-backend/repo/auth/auth_impl"
	conversation_repo "chatapp-backend/repo/conversation/conversation_impl"
	auth_service "chatapp-backend/service/auth/auth_impl"
	conversation_service "chatapp-backend/service/conversation/conversation_impl"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine, db *sql.DB) {

	// Auth Wiring

	authRepo := auth_repo.NewAuthImpl(db)
	authSvc := auth_service.NewAuth(authRepo)
	authHandler := auth_handler.NewHandler(authSvc)

	// Conversation Wiring

	conversationRepo := conversation_repo.NewConversationImpl(db)
	conversationSvc := conversation_service.NewConversation(conversationRepo)
	conversationHandler := conversation_handler.NewHandler(conversationSvc)

	//Auth Routes
	router.POST("/signup", authHandler.SignUP)
	router.POST("/login", authHandler.Login)

	// Conversation Routes

	protected := router.Group("/")
	protected.Use(middleware.AuthenticationMiddleware)
	{
		protected.POST("/dms", conversationHandler.StartDM)
		protected.GET("/dms", conversationHandler.ListDMs)
		protected.POST("/dms/:id/messages", conversationHandler.SendMessage)
		protected.GET("/dms/:id/messages", conversationHandler.ListMessages)
	}
}
