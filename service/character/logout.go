package character

import (
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/service"

	"gopkg.in/validator.v2"
)

type LogoutForm struct {
	Character *model.Character `validate:"nonzero"`
	Websocket service.Websocket
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

	payload := model.CharacterLogout {
		ID: self.Character.ID,
	}

	data := encapsulateTopic("logout", payload)
	self.Websocket.SendBroadcast(data)

	return nil, nil
}
