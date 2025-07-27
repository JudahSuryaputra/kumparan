package article

import (
	"context"
	"kumparan/internal/shared/dto"
)

func (s *implArticle) Get(ctx context.Context, filter *dto.ArticleFilter) (resp []dto.GetArticlesResponse, err error) {
	articles, err := s.repo.ArticleRepository.Get(filter)
	if err != nil {
		//TODO: LOG ERROR
		return nil, err
	}

	return articles, nil
}
