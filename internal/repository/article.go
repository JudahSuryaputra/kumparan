package repository

import (
	"fmt"
	"kumparan/internal/entity"
	"kumparan/internal/shared"
	"kumparan/internal/shared/dto"
)

type (
	ArticleRepository interface {
		Create(article *entity.Article) (err error)
		Get(filter *dto.ArticleFilter) (result []dto.GetArticlesResponse, err error)
	}

	implArticle struct {
		deps shared.Deps
	}
)

func NewArticleRepository(deps shared.Deps) ArticleRepository {
	return &implArticle{deps: deps}
}

func (r *implArticle) Create(article *entity.Article) (err error) {
	_, err = r.deps.DB.Exec(`
			INSERT INTO articles (id, author_id, title, body, created_at)
			VALUES ($1, $2, $3, $4, NOW())
		`, article.ID, article.AuthorID, article.Title, article.Body)

	if err != nil {
		return err
	}

	return
}

func (r *implArticle) Get(filter *dto.ArticleFilter) (result []dto.GetArticlesResponse, err error) {
	baseQuery := `
			SELECT articles.id, authors.name, articles.title, articles.body, articles.created_at
			FROM articles articles
			JOIN authors authors ON articles.author_id = authors.id
			WHERE 1=1
		`
	args := []interface{}{}
	argIdx := 1
	if filter.Query != "" {
		baseQuery += fmt.Sprintf(" AND (articles.title ILIKE $%d OR articles.body ILIKE $%d)", argIdx, argIdx+1)
		args = append(args, "%"+filter.Query+"%", "%"+filter.Query+"%")
		argIdx += 2
	}
	if filter.Author != "" {
		baseQuery += fmt.Sprintf(" AND authors.name ILIKE $%d", argIdx)
		args = append(args, "%"+filter.Author+"%")
		argIdx++
	}

	baseQuery += fmt.Sprintf(" ORDER BY articles.created_at DESC LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, filter.Limit, filter.Offset)

	rows, err := r.deps.DB.Query(baseQuery, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var article dto.GetArticlesResponse
		if err = rows.Scan(&article.ID, &article.AuthorName, &article.Title, &article.Body, &article.CreatedAt); err != nil {
			return result, err
		}
		result = append(result, article)
	}
	return
}
