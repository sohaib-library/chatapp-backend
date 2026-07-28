package ws

import (
	"encoding/json"
	"sync"

	"chatapp-backend/models"
)

type Event struct {
	Type           string          `json:"type"`
	UserID         string          `json:"user_id,omitempty"`
	Status         string          `json:"status,omitempty"`
	ConversationID string          `json:"conversation_id,omitempty"`
	Message        *models.Message `json:"message,omitempty"`
	OnlineUsers    []string        `json:"online_users,omitempty"`
}

type Hub struct {
	mu         sync.RWMutex
	clients    map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if h.clients[client.UserID] == nil {
				h.clients[client.UserID] = make(map[*Client]bool)
			}
			wasOffline := len(h.clients[client.UserID]) == 0
			h.clients[client.UserID][client] = true
			h.mu.Unlock()

			if wasOffline {
				h.BroadcastPresence(client.UserID, "online")
			}
			h.sendOnlineUsers(client)

		case client := <-h.unregister:
			h.mu.Lock()
			if conns, ok := h.clients[client.UserID]; ok {
				if _, exists := conns[client]; exists {
					delete(conns, client)
					close(client.send)
					becameOffline := len(conns) == 0
					if becameOffline {
						delete(h.clients, client.UserID)
					}
					h.mu.Unlock()
					if becameOffline {
						h.BroadcastPresence(client.UserID, "offline")
					}
					continue
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			for _, conns := range h.clients {
				for client := range conns {
					select {
					case client.send <- message:
					default:
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *Hub) OnlineUserIDs() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	ids := make([]string, 0, len(h.clients))
	for userID := range h.clients {
		ids = append(ids, userID)
	}
	return ids
}

func (h *Hub) IsOnline(userID string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients[userID]) > 0
}

func (h *Hub) BroadcastPresence(userID, status string) {
	payload, err := json.Marshal(Event{
		Type:   "presence",
		UserID: userID,
		Status: status,
	})
	if err != nil {
		return
	}
	h.broadcast <- payload
}

func (h *Hub) NotifyMessage(conversationID string, memberIDs []string, message *models.Message) {
	payload, err := json.Marshal(Event{
		Type:           "message",
		ConversationID: conversationID,
		Message:        message,
	})
	if err != nil {
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, memberID := range memberIDs {
		conns, ok := h.clients[memberID]
		if !ok {
			continue
		}
		for client := range conns {
			select {
			case client.send <- payload:
			default:
			}
		}
	}
}

func (h *Hub) sendOnlineUsers(client *Client) {
	payload, err := json.Marshal(Event{
		Type:        "online_users",
		OnlineUsers: h.OnlineUserIDs(),
	})
	if err != nil {
		return
	}

	select {
	case client.send <- payload:
	default:
	}
}
