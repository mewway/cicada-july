package provider

import (
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/contracts/queue"
	"github.com/mewway/go-laravel/facades"
)

type QueueServiceProvider struct {
}

func (receiver *QueueServiceProvider) Register(app foundation.Application) {
	facades.Queue().Register(receiver.Jobs())
}

func (receiver *QueueServiceProvider) Boot(app foundation.Application) {

}

func (receiver *QueueServiceProvider) Jobs() []queue.Job {
	return []queue.Job{}
}
