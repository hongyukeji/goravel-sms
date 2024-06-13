package sms

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/hongyukeji/goravel-sms/commands"
)

const Binding = "sms"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/hongyukeji/goravel-sms", map[string]string{
		"config/sms.go": app.ConfigPath("sms.go"),
	})
	app.Commands([]console.Command{
		commands.NewHello(),
	})
}
