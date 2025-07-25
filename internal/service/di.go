package service

import (
	"go.uber.org/dig"
	"kumparan/internal/service/article"
	"kumparan/internal/service/platform"
)

type (
	Holder struct {
		dig.In

		PlatformService platform.Service
		Article         article.Service
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(platform.NewPlatformService); err != nil {
		return err
	}
	if err := container.Provide(article.NewArticleService); err != nil {
		return err
	}
	return nil
}
