package user

import "go-lemon/config"

type Module struct {
	Config config.Configuration
}

func InitModule(conf config.Configuration) *Module {
	return &Module{
		Config: conf,
	}
}
