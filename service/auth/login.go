package auth

import (
	"github.com/danielbintar/qwe-server/model"

	"github.com/danielbintar/go-record/db"

	"gopkg.in/validator.v2"
)

type LoginForm struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func (self *LoginForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *LoginForm) Perform() (interface{}, []error) {
	user := &model.User{}
	err := db.FindBy(&user, []string{"username", "=", self.Username}, []string{"password", "=", self.Password})

	if err != nil {
		return user, []error{err}
	}

	return user, nil
}
