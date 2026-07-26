package route

import (
	"database/sql"

	auth_handler "chatapp-backend/handler/auth/auth_impl"
	conversation_handler "chatapp-backend/handler/conversation/conversation_impl"
	user_handler "chatapp-backend/handler/user/user_impl"
	"chatapp-backend/middleware"
	auth_repo "chatapp-backend/repo/auth/auth_impl"
	conversation_repo "chatapp-backend/repo/conversation/conversation_impl"
	user_repo "chatapp-backend/repo/user/user_impl"
	auth_service "chatapp-backend/service/auth/auth_impl"
	conversation_service "chatapp-backend/service/conversation/conversation_impl"
	user_service "chatapp-backend/service/user/user_impl"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine, db *sql.DB) {
 
	// Auth Wiring
	authRepo := auth_repo.NewAuthImpl(db)
	authSvc := auth_service.NewAuth(authRepo)
	authHandler := auth_handler.NewHandler(authSvc)

	// User Wiring

	userRepo := user_repo.NewUserImpl(db)
	userSvc := user_service.NewUser(userRepo)
	userHandler := user_handler.NewHandler(userSvc)

	// Conversation Wiring

	conversationRepo := conversation_repo.NewConversationImpl(db)
	conversationSvc := conversation_service.NewConversation(conversationRepo)
	conversationHandler := conversation_handler.NewHandler(conversationSvc)

	// Auth Routes

	router.POST("/signup", authHandler.SignUp)
	router.POST("/login", authHandler.Login)

	

	// Routes	

	protected := router.Group("/")
	protected.Use(middleware.AuthenticationMiddleware)
	{
		// User Routes
		
		protected.GET("/users", userHandler.ListUsers)

		// Conversation Routes

		protected.POST("/dms", conversationHandler.StartDM)
		protected.GET("/dms", conversationHandler.ListDMs)



		protected.POST("/groups", conversationHandler.CreateGroup)
		protected.GET("/groups", conversationHandler.ListGroups)

		// Messages work for both DMs and groups (same conversation + membership check)
		
		protected.POST("/dms/:id/messages", conversationHandler.SendMessage)
		protected.GET("/dms/:id/messages", conversationHandler.ListMessages)
		protected.POST("/groups/:id/messages", conversationHandler.SendMessage)
		protected.GET("/groups/:id/messages", conversationHandler.ListMessages)
	}
}
