package article

import (
	"context"
	"kumparan/internal/repository"
	"kumparan/internal/shared"
	"kumparan/internal/shared/dto"
)

type (
	Service interface {
		Create(ctx context.Context, createRequest *dto.CreateArticleRequest) (err error)
		Get(ctx context.Context, filter *dto.ArticleFilter) (resp []dto.GetArticlesResponse, err error)
	}

	implArticle struct {
		deps shared.Deps

		repo repository.Holder
	}
)

func NewArticleService(deps shared.Deps, repo repository.Holder) Service {
	return &implArticle{deps: deps, repo: repo}
}
