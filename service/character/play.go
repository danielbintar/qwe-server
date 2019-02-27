package character

import (
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/service"

	"gopkg.in/validator.v2"
)

type PlayForm struct {
	Character *model.Character `validate:"nonzero"`
	Websocket service.Websocket
}

func (self *PlayForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *PlayForm) Perform() (interface{}, []error) {
	place := repository.GetCharacterActivePlace(self.Character.ID)
	if place == nil {
		defaultPlace := "town"
		place = &defaultPlace
		repository.SetCharacterActivePlace(self.Character.ID, *place)
	}
	self.Character.ActivePlace = place

	repository.SetLoginCharacter(self.Character.ID)
	repository.SetCurrentCharacter(self.Character.UserID, self.Character.ID)

	switch *place {
	case "town":
		self.ManageTown()
	case "region":
		self.ManageRegion()
	}

	return self.Character, nil
}

func (self *PlayForm) ManageTown() {
	townID := repository.GetCharacterTownID(self.Character.ID)
	if townID == nil {
		defaultTownID := uint(1)
		townID = &defaultTownID
		repository.SetCharacterTownID(self.Character.ID, defaultTownID)
	}

	position := repository.GetTownCharacterPosition(*townID, self.Character.ID)
	if position == nil {
		town := repository.FindTown(*townID)

		position = &model.CharacterPosition {
			ID: self.Character.ID,
			X: town.Position.X,
			Y: town.Position.Y,
		}

		repository.SetTownCharacterPosition(*townID, *position)
	}

	resp := model.MoveOutgoing {
		X: position.X,
		Y: position.Y,
		ActivePlace: "town",
		Character: model.MoveCharacter {
			ID: position.ID,
		},
	}

	data := encapsulateTopic("move", resp)
	self.Websocket.SendBroadcast(data)
}

func (self *PlayForm) ManageRegion() {
	regionID := repository.GetCharacterRegionID(self.Character.ID)
	position := repository.GetRegionCharacterPosition(*regionID, self.Character.ID)

	resp := model.MoveOutgoing {
		X: position.X,
		Y: position.Y,
		ActivePlace: "region",
		Character: model.MoveCharacter {
			ID: position.ID,
		},
	}

	data := encapsulateTopic("move", resp)
	self.Websocket.SendBroadcast(data)
}

func encapsulateTopic(action string, data interface{}) []byte {
	o := model.OutgoingMessage {
		Action: action,
		Data: data,
	}
	b, _ := json.Marshal(&o)
	return b
}
