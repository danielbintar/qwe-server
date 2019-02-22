package websocket

import (
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
)

func (c Client) manageChat(rawData []byte) {
	var req model.ChatIncoming
	err := json.Unmarshal(rawData, &req)
	if err != nil { return }

	resp := model.ChatOutgoing {
		Message: req.Message,
		Sender: model.ChatSender {
			ID: c.character.ID,
			Name: c.character.Name,
		},
	}

	data := encapsulateTopic("chat", resp)
	c.hub.Broadcast <- data
}
