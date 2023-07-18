package provider

import (
	"cicada-july/app/middleware"
	"cicada-july/route"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/facades"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
	// Add HTTP middlewares
	facades.Route().GlobalMiddleware(middleware.Kernel{}.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	receiver.configureRateLimiting()
	// 注册前台路由
	route.Api()
	// 注册后台路由
	route.Admin()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {

}
