package user

import (
	"errors"

	"github.com/danielbintar/qwe-server/model"

	"github.com/danielbintar/qwe-server/db"

	"gopkg.in/validator.v2"

	"golang.org/x/crypto/bcrypt"
)

type CreateForm struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func (self *CreateForm) Validate() []error {
	var errs []error

	if err := validator.Validate(self); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (self *CreateForm) Perform() (interface{}, []error) {
	user := &model.User{Username: self.Username}

	if !db.DB().Where(&user).First(&user).RecordNotFound() {
		return user, []error{errors.New("username already used")}
	}

	// 0 for using default cost
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(self.Password), 0)
	user.Username = self.Username
	user.Password = string(encryptedPassword)
	db.DB().Create(&user)

	return user, nil
}
