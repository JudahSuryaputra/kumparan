package dto

import "time"

type CreateArticleRequest struct {
	AuthorID string `json:"author_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
}

type ArticleFilter struct {
	Query  string
	Author string
	Limit  int
	Offset int
}

type GetArticlesResponse struct {
	ID         string    `json:"id"`
	AuthorName string    `json:"author_name"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
}
