package auth

import (
	"errors"

	"github.com/danielbintar/qwe-server/model"

	"github.com/danielbintar/qwe-server/db"

	"gopkg.in/validator.v2"

	"golang.org/x/crypto/bcrypt"
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
	user := &model.User{Username: self.Username}
	db.DB().Where(&user).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(self.Password)); err != nil {
		return user, []error{errors.New("not found")}
	}

	return user, nil
}
