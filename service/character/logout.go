package character

import (
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"

	"gopkg.in/validator.v2"
)

type LogoutForm struct {
	Character *model.Character `validate:"nonzero"`
}

func (self *LogoutForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *LogoutForm) Perform() (interface{}, []error) {
	repository.UnsetLoginCharacter(self.Character.ID)

	return nil, nil
}
