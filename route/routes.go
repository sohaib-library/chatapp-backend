package route

import (
	auth_handler "chatapp-backend/handler/auth/auth_impl"
	conversation_handler "chatapp-backend/handler/conversation/conversation_impl"
	user_handler "chatapp-backend/handler/user/user_impl"
	"chatapp-backend/middleware"
	"chatapp-backend/mqtt"
	auth_repo "chatapp-backend/repo/auth/auth_impl"
	conversation_repo "chatapp-backend/repo/conversation/conversation_impl"
	user_repo "chatapp-backend/repo/user/user_impl"
	auth_service "chatapp-backend/service/auth/auth_impl"
	conversation_service "chatapp-backend/service/conversation/conversation_impl"
	user_service "chatapp-backend/service/user/user_impl"
	"chatapp-backend/ws"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoute(router *gin.Engine, db *gorm.DB) {

	// MQTT notifier — publishes messages to EMQX for real-time delivery.
	mqttClient := mqtt.NewClient()
	notifier := mqtt.NewMQTTNotifier(mqttClient)

	// WebSocket hub — kept for browser clients that connect via /ws.
	hub := ws.NewHub()
	go hub.Run()

	authRepo := auth_repo.NewAuthImpl(db)
	authSvc := auth_service.NewAuth(authRepo)
	authHandler := auth_handler.NewHandler(authSvc)

	// User Wiring

	userRepo := user_repo.NewUserImpl(db)
	userSvc := user_service.NewUser(userRepo)
	userHandler := user_handler.NewHandler(userSvc)

	// Conversation Wiring — uses MQTT notifier for real-time message events.

	conversationRepo := conversation_repo.NewConversationImpl(db)
	conversationSvc := conversation_service.NewConversation(conversationRepo, notifier)
	conversationHandler := conversation_handler.NewHandler(conversationSvc)

	// Health check — used by K8s readiness/liveness probes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Auth Routes

	router.POST("/signup", authHandler.SignUp)
	router.POST("/login", authHandler.Login)

	// WebSocket endpoint (browser clients can still connect here)
	router.GET("/ws", ws.ServeWS(hub))

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
