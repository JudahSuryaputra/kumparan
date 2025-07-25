package article

import (
	"context"
	"kumparan/internal/controller/helper"
	"kumparan/internal/entity"
	"kumparan/internal/shared/dto"
	"time"
)

func (s *implArticle) Create(ctx context.Context, createRequest *dto.CreateArticleRequest) (err error) {
	article := entity.Article{
		ID:        helper.GenerateUUID(),
		AuthorID:  createRequest.AuthorID,
		Title:     createRequest.Title,
		Body:      createRequest.Body,
		CreatedAt: time.Now(),
	}

	err = s.repo.ArticleRepository.Create(&article)
	if err != nil {
		//TODO: LOG ERROR
		return err
	}
	return
}
