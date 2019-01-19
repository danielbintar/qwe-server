package auth

import (
	"github.com/danielbintar/qwe-server/service"
)

var Login = func(form LoginForm) (interface{}, []error) {
	return service.Start(&form)
}
