package middleware

import (
	"net/http"
	"strings"

	"chatapp-backend/utils"

	"github.com/gin-gonic/gin"
)

const UserIDKey = "userID"

func AuthenticationMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID, err := utils.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	ctx.Set(UserIDKey, userID)
	ctx.Next()
}

func GetUserID(ctx *gin.Context) (string, bool) {
	value, exists := ctx.Get(UserIDKey)
	if !exists {
		return "", false
	}

	userID, ok := value.(string)
	return userID, ok && userID != ""
}
