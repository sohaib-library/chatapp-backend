package conversation

import "chatapp-backend/models"

type RealtimeNotifier interface {
	NotifyMessage(conversationID string, memberIDs []string, message *models.Message)
}
