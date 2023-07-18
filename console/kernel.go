package console

import (
	"github.com/mewway/go-laravel/contracts/console"
	"github.com/mewway/go-laravel/contracts/schedule"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{}
}

func (kernel *Kernel) InitRoutes() {

}
