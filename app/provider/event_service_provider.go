package provider

import (
	"github.com/mewway/go-laravel/contracts/event"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/facades"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Register(app foundation.Application) {
	facades.Event().Register(receiver.listen())
}

func (receiver *EventServiceProvider) Boot(app foundation.Application) {

}

func (receiver *EventServiceProvider) listen() map[event.Event][]event.Listener {
	return map[event.Event][]event.Listener{}
}
