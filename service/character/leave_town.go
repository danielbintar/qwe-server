package character

import (
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/model"

	"gopkg.in/validator.v2"
)

type LeaveTownForm struct {
	ID uint `validate:"nonzero"`
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
	townID := repository.GetCharacterInTown(self.ID)
	if townID == nil { panic("panic") }

	repository.UnsetCharacterInTown(self.ID)
	repository.UnsetCharacterTownPosition(self.ID, *townID)

	town := repository.FindTown(*townID)
	region := repository.FindRegion(town.RegionID)
	townPosition := region.FindTownPosition(town.ID)
	position := &model.CharacterPosition {
		ID: self.ID,
		X: townPosition.X,
		Y: townPosition.Y,
	}
	repository.SetCharacterInRegion(self.ID, region.ID)
	repository.SetRegionCharacterPosition(region.ID, position)

	return nil, nil
}
