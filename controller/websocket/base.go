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
	characterService "github.com/danielbintar/qwe-server/service/character"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func encapsulateTopic(action string, data interface{}) []byte {
	o := model.OutgoingMessage {
		Action: action,
		Data: data,
	}
	b, _ := json.Marshal(&o)
	return b
}

func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()

		form := characterService.LogoutForm{
			Character: c.character,
			Websocket: c.hub,
		}
		characterService.Logout(form)
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
		var r *model.WebsocketRequest
		err = json.Unmarshal(message, &r)
		if err != nil { continue }

		switch r.Action {
		case "move":
			c.manageMove(r.Data)
		case "chat":
			c.manageChat(r.Data)
		}
	}
}

func (c *Client) write() {
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

			// Add queued messages to the current websocket message.
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

func Manage(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil { panic(err) }

	ctx := r.Context()
	userID := ctx.Value("jwt").(*model.Jwt).UserID
	characterID := repository.GetCurrentCharacter(userID)
	character := &model.Character{ID: *characterID}
	db.DB().Where(&character).First(&character)

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), character: character}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.write()
	go client.read()
}
