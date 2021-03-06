package service

type Service interface {
	Validate() []error
	Perform() (interface{}, []error)
}

type Websocket interface {
	SendBroadcast([]byte)
}

func Start(svc Service) (interface{}, []error) {
	var object interface{}

	errors := svc.Validate()
	if errors != nil {
		return object, errors
	}

	object, errors = svc.Perform()
	return object, errors
}
