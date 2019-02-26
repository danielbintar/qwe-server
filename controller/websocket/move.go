package websocket

import (
	"encoding/json"

	"github.com/danielbintar/qwe-server/constant"
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	characterService "github.com/danielbintar/qwe-server/service/character"
)

func (c Client) manageMove(rawData []byte) {
	var req model.MoveIncoming
	err := json.Unmarshal(rawData, &req)
	if err != nil { return }

	if req.Direction != "left" && req.Direction != "right" && req.Direction != "up" && req.Direction != "down" {
		return
	}

	place := repository.GetCharacterActivePlace(c.character.ID)
	switch *place {
	case "region":
		c.manageMoveRegion(req)
	case "town":
		c.manageMoveTown(req)
	default:
		c.send <- []byte(constant.PING)
	}


}

func (c Client) manageMoveTown(req model.MoveIncoming) {
	townID := repository.GetCharacterTownID(c.character.ID)
	if townID == nil { return }

	position := repository.GetTownCharacterPosition(*townID, c.character.ID)

	switch req.Direction {
	case "left":
		if position.X == 0 {
			c.send <- []byte(constant.PING)
			return
		} else {
			position.X--
		}
	case "right":
		position.X++
	case "down":
		if position.Y == 0 {
			c.send <- []byte(constant.PING)
			return
		} else {
			position.Y--
		}
	case "up":
		position.Y++
	}

	town := repository.FindTown(*townID)
	for _, portal := range town.Portals {
		if portal.In(*position) {
			form := characterService.LeaveTownForm {
				Character: *c.character,
				Websocket: c.hub,
			}
			characterService.LeaveTown(form)

			return
		}
	}


	repository.SetTownCharacterPosition(*townID, *position)

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

func (c Client) manageMoveRegion(req model.MoveIncoming) {
	regionID := repository.GetCharacterRegionID(c.character.ID)
	if regionID == nil { return }

	position := repository.GetRegionCharacterPosition(*regionID, c.character.ID)

	switch req.Direction {
	case "left":
		if position.X == 0 {
			c.send <- []byte(constant.PING)
			return
		} else {
			position.X--
		}
	case "right":
		position.X++
	case "down":
		if position.Y == 0 {
			c.send <- []byte(constant.PING)
			return
		} else {
			position.Y--
		}
	case "up":
		position.Y++
	}

	region := repository.FindRegion(*regionID)
	for _, town := range region.Towns {
		if town.Portal.In(*position) {
			form := characterService.EnterTownForm {
				Character: *c.character,
				TownID: town.ID,
				Websocket: c.hub,
			}
			characterService.EnterTown(form)

			return
		}
	}

	repository.SetRegionCharacterPosition(*regionID, *position)

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
