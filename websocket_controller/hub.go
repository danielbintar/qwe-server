package websocket_controller

import "sync"

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

var (
	chatHub *Hub
	chatHubMutex sync.Once

	moveHub *Hub
	moveHubMutex sync.Once
)

func ChatHubInstance() *Hub {
	chatHubMutex.Do(func() {
		chatHub = newHub()
	})
	return chatHub
}

func MoveHubInstance() *Hub {
	moveHubMutex.Do(func() {
		moveHub = newHub()
	})
	return moveHub
}

func newHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
