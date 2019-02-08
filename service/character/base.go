package character

import "github.com/danielbintar/qwe-server/service"

var Create = func(form CreateForm) (interface{}, []error) {
	return service.Start(&form)
}
