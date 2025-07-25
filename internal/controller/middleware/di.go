package middleware

import (
	"go.uber.org/dig"
	"kumparan/internal/repository"
	"kumparan/internal/shared"
)

type Middleware interface {
}

type implMiddleware struct {
	repo repository.Holder
	deps shared.Deps
}

func NewMiddlewares(repo repository.Holder, deps shared.Deps) Middleware {
	return &implMiddleware{repo: repo, deps: deps}
}

func Register(container *dig.Container) error {
	return container.Provide(NewMiddlewares)
}
