package main

import (
	"cicada-july/bootstrap"
	"github.com/mewway/go-laravel/facades"
)

func main() {
	bootstrap.Boot()
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
