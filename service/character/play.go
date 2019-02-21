package character

import (
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/websocket_controller"

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
	position := repository.SetCurrentCharacter(self.Character.UserID, self.Character.ID)
	repository.SetLoginCharacter(self.Character.ID)

	encodedPosition, _ := json.Marshal(position)
	websocket_controller.MoveHubInstance().Broadcast <- []byte(encodedPosition)

	return nil, nil
}
