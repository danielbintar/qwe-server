package websocket

import (
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
)

func (c Client) manageMove(rawData []byte) {
	var req model.MoveIncoming
	err := json.Unmarshal(rawData, &req)
	if err != nil { return }

	if req.Direction != "left" && req.Direction != "right" && req.Direction != "up" && req.Direction != "down" {
		return
	}

	townID := repository.GetCharacterTownID(c.character.ID)
	if townID == nil { return }

	position := repository.GetTownCharacterPosition(*townID, c.character.ID)

	switch req.Direction {
	case "left":
		if position.X == 0 {
			return
		} else {
			position.X--
		}
	case "right":
		position.X++
	case "down":
		if position.Y == 0 {
			return
		} else {
			position.Y--
		}
	case "up":
		position.Y++
	}

	repository.SetTownCharacterPosition(*townID, position)

	resp := model.MoveOutgoing {
		X: position.X,
		Y: position.Y,
		Character: model.MoveCharacter {
			ID: c.character.ID,
		},
	}

	data := encapsulateTopic("move", resp)
	c.hub.Broadcast <- data
}
