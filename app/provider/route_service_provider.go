package provider

import (
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/facades"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
	// Add HTTP middlewares
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	receiver.configureRateLimiting()

	routes.Web()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {

}
