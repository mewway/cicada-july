package config

import (
	"cicada-july/app/provider"
	"github.com/mewway/go-laravel/auth"
	"github.com/mewway/go-laravel/cache"
	"github.com/mewway/go-laravel/console"
	"github.com/mewway/go-laravel/contracts/foundation"
	"github.com/mewway/go-laravel/crypt"
	"github.com/mewway/go-laravel/database"
	"github.com/mewway/go-laravel/event"
	"github.com/mewway/go-laravel/facades"
	"github.com/mewway/go-laravel/filesystem"
	"github.com/mewway/go-laravel/grpc"
	"github.com/mewway/go-laravel/hash"
	"github.com/mewway/go-laravel/http"
	"github.com/mewway/go-laravel/log"
	"github.com/mewway/go-laravel/mail"
	"github.com/mewway/go-laravel/queue"
	"github.com/mewway/go-laravel/route"
	"github.com/mewway/go-laravel/schedule"
	"github.com/mewway/go-laravel/support/carbon"
	"github.com/mewway/go-laravel/validation"
)

func Boot() {

}

func init() {
	config := facades.Config()

	config.Add("app", map[string]any{
		// Application Name
		//
		// This value is the name of your application. This value is used when the
		// framework needs to place the application's name in a notification or
		// any other location as required by the application or its packages.
		"name": config.Env("APP_NAME", "Goravel"),

		// Application Environment
		//
		// This value determines the "environment" your application is currently
		// running in. This may determine how you prefer to configure various
		// services the application utilizes. Set this in your ".env" file.
		"env": config.Env("APP_ENV", "production"),

		// Application Debug Mode
		"debug": config.Env("APP_DEBUG", false),

		// Application Timezone
		//
		// Here you may specify the default timezone for your application.
		// Example: UTC, Asia/Shanghai
		// More: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
		"timezone": carbon.UTC,

		// Encryption Key
		//
		// 32 character string, otherwise these encrypted strings
		// will not be safe. Please do this before deploying an application!
		"key": config.Env("APP_KEY", ""),

		// Autoload service providers
		//
		// The service providers listed here will be automatically loaded on the
		// request to your application. Feel free to add your own services to
		// this array to grant expanded functionality to your applications.
		"providers": []foundation.ServiceProvider{
			&log.ServiceProvider{},
			&console.ServiceProvider{},
			&database.ServiceProvider{},
			&cache.ServiceProvider{},
			&http.ServiceProvider{},
			&route.ServiceProvider{},
			&schedule.ServiceProvider{},
			&event.ServiceProvider{},
			&queue.ServiceProvider{},
			&grpc.ServiceProvider{},
			&mail.ServiceProvider{},
			&auth.ServiceProvider{},
			&hash.ServiceProvider{},
			&crypt.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&validation.ServiceProvider{},
			&provider.AppServiceProvider{},
			&provider.AuthServiceProvider{},
			&provider.RouteServiceProvider{},
			&provider.GrpcServiceProvider{},
			&provider.ConsoleServiceProvider{},
			&provider.QueueServiceProvider{},
			&provider.EventServiceProvider{},
			&provider.ValidationServiceProvider{},
		},
	})
}
