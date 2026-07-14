package route

import (
	"database/sql"

	auth_handler "chatapp-backend/handler/auth/auth_impl"
	auth_repo "chatapp-backend/repo/auth/auth_impl"
	auth_service "chatapp-backend/service/auth/auth_impl"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine, db *sql.DB) {

	// Pipeline
	// Auth_Signup
	authRepo := auth_repo.NewAuthImpl(db)
	authSvc := auth_service.NewAuth(authRepo)
	h := auth_handler.NewHandler(authSvc)


   // Routes

	// Auth
	router.POST("/signup", h.SignUP)
	router.POST("/login", h.Login)



}
