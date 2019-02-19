package character

import (
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"

	"gopkg.in/validator.v2"
)

type PlayForm struct {
	Character *model.Character `validate:"nonzero"`
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
	repository.SetCurrentCharacter(self.Character.UserID, self.Character.ID)
	repository.SetLoginCharacter(self.Character.ID)

	return nil, nil
}
