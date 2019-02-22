package websocket

import (
	"bytes"
	"log"
	"net/http"
	"time"
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/db"
	"github.com/danielbintar/qwe-server/repository"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *Client) readChat() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var chat *model.Chat
		json.Unmarshal(message, &chat)
		chat.Sender = c.character.Name
		encodedChat, _ := json.Marshal(chat)
		c.hub.Broadcast <- []byte(string(encodedChat))
	}
}

func (c *Client) writeChat() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ManageChat(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil { panic(err) }

	ctx := r.Context()
	userID := ctx.Value("jwt").(*model.Jwt).UserID
	characterID := repository.GetCurrentCharacter(userID)
	character := &model.Character{ID: *characterID}
	db.DB().Where(&character).First(&character)

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), character: *character}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writeChat()
	go client.readChat()
}
