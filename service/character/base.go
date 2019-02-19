package character

import "github.com/danielbintar/qwe-server/service"

var Create = func(form CreateForm) (interface{}, []error) {
	return service.Start(&form)
}

var Play = func(form PlayForm) (interface{}, []error) {
	return service.Start(&form)
}

var Logout = func(form LogoutForm) (interface{}, []error) {
	return service.Start(&form)
}
