package di

import (
	"go.uber.org/dig"
	"kumparan/config"
	"kumparan/internal/controller"
	"kumparan/internal/repository"
	"kumparan/internal/service"
	"kumparan/internal/shared"
)

var (
	Container = dig.New()
)

func init() {
	if err := Container.Provide(config.New); err != nil {
		panic(err)
	}
	if err := Container.Provide(NewORM); err != nil {
		panic(err)
	}
	if err := Container.Provide(NewRedisClient); err != nil {
		panic(err)
	}

	if err := controller.Register(Container); err != nil {
		panic(err)
	}
	if err := service.Register(Container); err != nil {
		panic(err)
	}
	if err := repository.Register(Container); err != nil {
		panic(err)
	}
	if err := shared.Register(Container); err != nil {
		panic(err)
	}
}
