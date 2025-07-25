package dto

import "time"

type CreateArticleRequest struct {
	AuthorID string `json:"author_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
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
