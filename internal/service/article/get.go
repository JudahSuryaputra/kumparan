package article

import (
	"context"
	"kumparan/internal/controller/helper"
	"kumparan/internal/entity"
	"kumparan/internal/shared/dto"
	"strconv"
	"time"
)

func (s *implArticle) Get(ctx context.Context, filter *dto.ArticleFilter) (err error) {
	articles, err := s.repo.ArticleRepository.Get(filter)
	if err != nil {
		//TODO: LOG ERROR
		return err
	}

	return articles, nil
}
