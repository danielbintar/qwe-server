package character

import (
	"github.com/danielbintar/qwe-server/repository"

	"gopkg.in/validator.v2"
)

type EnterTownFromRegionForm struct {
	ID     uint `validate:"nonzero"`
	TownID uint `validate:"nonzero"`
}

func (self *EnterTownFromRegionForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *EnterTownFromRegionForm) Perform() (interface{}, []error) {
	repository.UnsetCharacterInRegion(self.ID)
	repository.UnsetRegionCharacterPosition(self.ID, self.TownID)

	form := EnterTownForm {
		ID: self.ID,
		TownID: self.TownID,
	}

	EnterTown(form)

	return nil, nil
}
