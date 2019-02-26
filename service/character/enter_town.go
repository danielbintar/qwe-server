package character

import (
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/service"

	"gopkg.in/validator.v2"
)

type EnterTownForm struct {
	Character model.Character   `validate:"nonzero"`
	TownID    uint              `validate:"nonzero"`
	Websocket service.Websocket
}

func (self *EnterTownForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *EnterTownForm) Perform() (interface{}, []error) {
	regionID := repository.GetCharacterRegionID(self.Character.ID)
	if regionID == nil {
		return nil, nil
	}

	repository.UnsetRegionCharacterPosition(*regionID, self.Character.ID)
	repository.UnsetCharacterRegionID(self.Character.ID)

	town := repository.FindTown(self.TownID)
	repository.SetCharacterTownID(self.Character.ID, town.ID)
	repository.SetCharacterActivePlace(self.Character.ID, "town")

	position := model.CharacterPosition {
		ID: self.Character.ID,
		X: town.Position.X,
		Y: town.Position.Y,
	}

	repository.SetTownCharacterPosition(town.ID, position)

	resp := model.LeaveRegionData {ID: self.Character.ID}
	data := encapsulateTopic("leave_region", resp)
	self.Websocket.SendBroadcast(data)

	return nil, nil
}
