package pkg

import "github.com/go-kit/log"

type Authentication interface {
	Authenticate(username, password string) (interface{}, error)
}

type authentication struct {
	logger *log.Logger
}

func NewAuthenticationSvc(logger *log.Logger) Authentication {
	return &authentication{
		logger: logger,
	}
}

func (a *authentication) Authenticate(username, password string) (interface{}, error) {
	return nil, nil
}
