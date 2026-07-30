package mqtt

import (
	"encoding/json"
	"fmt"
	"log"

	"chatapp-backend/models"
	conversationService "chatapp-backend/service/conversation"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// Ensure MQTTNotifier implements the RealtimeNotifier interface.
var _ conversationService.RealtimeNotifier = (*MQTTNotifier)(nil)

// MQTTNotifier publishes new messages to EMQX on a per-conversation topic.
// Topic pattern: chat/conversation/{conversationID}
//
// Clients (web, mobile, MQTTX for testing) subscribe to:
//   - chat/conversation/{id}  — specific conversation
//   - chat/conversation/+     — all conversations
type MQTTNotifier struct {
	client paho.Client
}

func NewMQTTNotifier(client paho.Client) *MQTTNotifier {
	return &MQTTNotifier{client: client}
}

type messageEvent struct {
	Type           string          `json:"type"`
	ConversationID string          `json:"conversation_id"`
	Message        *models.Message `json:"message"`
}

// NotifyMessage publishes the message to chat/conversation/{conversationID}.
// QoS 1 ensures at-least-once delivery. Retained=false so old messages
// don't replay on new subscriptions.
func (n *MQTTNotifier) NotifyMessage(conversationID string, _ []string, message *models.Message) {
	topic := fmt.Sprintf("chat/conversation/%s", conversationID)

	payload, err := json.Marshal(messageEvent{
		Type:           "message",
		ConversationID: conversationID,
		Message:        message,
	})
	if err != nil {
		log.Printf("[MQTT] marshal error: %v", err)
		return
	}

	token := n.client.Publish(topic, 1, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("[MQTT] publish error on topic %s: %v", topic, token.Error())
	}
}
