package middleware

import "github.com/mewway/go-laravel/contracts/http"

type Kernel struct {
}

func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{}
}
