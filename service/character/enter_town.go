package character

import (
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"

	"gopkg.in/validator.v2"
)

type EnterTownForm struct {
	ID     uint `validate:"nonzero"`
	TownID uint `validate:"nonzero"`
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
	town := repository.FindTown(self.TownID)
	position := model.CharacterPosition {
		ID: self.ID,
		X: town.Position.X,
		Y: town.Position.Y,
	}

	repository.SetCharacterInTown(self.ID, self.TownID)
	repository.SetTownCharacterPosition(self.TownID, &position)

	return nil, nil
}
