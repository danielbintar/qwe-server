package character

import (
	"github.com/danielbintar/qwe-server/repository"

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

	return nil, nil
}
