package bootstrap

import (
	"cicada-july/config"
	"github.com/mewway/go-laravel/foundation"
)

func Boot() {
	app := foundation.NewApplication()

	app.Boot()

	config.Boot()
}
