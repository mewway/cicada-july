package provider

import (
	"cicada-july/grpc"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/facades"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register(app foundation.Application) {
	// Add Grpc interceptors
	kernel := grpc.Kernel{}
	facades.Grpc().UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	facades.Grpc().UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot(app foundation.Application) {
	// Add routes
	grpc.InitRouters()
}
