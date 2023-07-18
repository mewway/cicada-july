package route

import (
	"github.com/mewway/go-laravel/contracts/http"
	"github.com/mewway/go-laravel/facades"
)

func Api() {
	facades.Route().Get("/", func(context http.Context) {
		context.Response().Json(http.StatusOK, http.Json{
			"hola": "me",
		})
	})
}

func Admin() {
	facades.Route().Get("/admin", func(context http.Context) {
		context.Response().Json(http.StatusOK, http.Json{
			"hola": "admin",
		})
	})
}
