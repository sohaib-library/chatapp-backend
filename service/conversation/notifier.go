package conversation

import "github.com/sohaib-library/chatapp-backend/models"

type RealtimeNotifier interface {
	NotifyMessage(conversationID string, memberIDs []string, message *models.Message)
}
