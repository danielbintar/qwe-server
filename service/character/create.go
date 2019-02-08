package character

import (
	"fmt"
	"errors"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/db"
	"github.com/danielbintar/qwe-server/constant"

	"gopkg.in/validator.v2"
)

type CreateForm struct {
	Name   string `json:"name" validate:"nonzero"`
	UserID uint   `json:"user_id" validate:"nonzero"`
}

func (self *CreateForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	if self.AlreadyAtLimit() {
		return []error{errors.New("total character exceed limit")}
	}

	if self.Exists() {
		return []error{errors.New("name already used")}
	}

	return nil
}

func (self *CreateForm) AlreadyAtLimit() bool {
	var charactersCount uint
	character := &model.Character{UserID: self.UserID}
	db.DB().Where(character).Count(&charactersCount)

	return charactersCount == constant.MY_CHARACTERS_LIMIT
}

func (self *CreateForm) Exists() bool {
	character := &model.Character{Name: self.Name}
	db.DB().Where(character).Take(character)
	return !db.DB().NewRecord(character)
}


func (self *CreateForm) Perform() (interface{}, []error) {
	character := &model.Character{UserID: self.UserID, Name: self.Name}
	db.DB().Create(character)
	return character, nil
}
