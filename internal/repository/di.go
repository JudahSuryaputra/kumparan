package repository

import (
	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	CacheRepository   CacheRepository
	ArticleRepository ArticleRepository
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewArticleRepository); err != nil {
		return err
	}
	if err := container.Provide(NewCacheRepository); err != nil {
		return err
	}
	return nil
}
