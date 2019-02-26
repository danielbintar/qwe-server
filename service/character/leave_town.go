package character

import (
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/service"

	"gopkg.in/validator.v2"
)

type LeaveTownForm struct {
	Character model.Character   `validate:"nonzero"`
	Websocket service.Websocket
}

func (self *LeaveTownForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *LeaveTownForm) Perform() (interface{}, []error) {
	townID := repository.GetCharacterTownID(self.Character.ID)
	if townID == nil {
		return nil, nil
	}

	repository.UnsetTownCharacterPosition(*townID, self.Character.ID)
	repository.UnsetCharacterTownID(self.Character.ID)

	town := repository.FindTown(*townID)
	region := repository.FindRegion(town.RegionID)
	repository.SetCharacterRegionID(self.Character.ID, region.ID)
	repository.SetCharacterActivePlace(self.Character.ID, "region")

	regionTown := region.FindTown(*townID)
	position := model.CharacterPosition {
		ID: self.Character.ID,
		X: regionTown.SpawnPosition.X,
		Y: regionTown.SpawnPosition.Y,
	}

	repository.SetRegionCharacterPosition(region.ID, position)

	resp := model.LeaveTownData {ID: self.Character.ID}
	data := encapsulateTopic("leave_town", resp)
	self.Websocket.SendBroadcast(data)

	return nil, nil
}
