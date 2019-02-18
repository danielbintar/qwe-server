package character

import (
	"github.com/danielbintar/qwe-server/repository"

	"gopkg.in/validator.v2"
)

type PlayForm struct {
	UserID      uint `validate:"nonzero"`
	CharacterID uint `validate:"nonzero"`
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
	repository.SetPlayingCharacter(self.UserID, self.CharacterID)
	townID := repository.GetCharacterInTown(self.CharacterID)

	if townID == nil {
		form := EnterTownForm {
			ID: self.CharacterID,
			TownID: 1,
		}

		EnterTown(form)
	}

	return nil, nil
}
