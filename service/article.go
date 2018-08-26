package service

import (
	"github.com/kongoole/minreuse-go/model"
)

type articleService struct{}

type AddArticleParams struct {
	Title string `json:"title"`
	Content string `json:"content"`
	TagIds []int `json:"tag_ids"`
}

func NewArticleService() *articleService {
	return new(articleService)
}

func (as *articleService) UpdateArticle(articleId int, data map[string]interface{}) bool {
	return model.ArticleModelInstance().UpdateArticle(articleId, data)
}

func (as *articleService) AddArticle(params AddArticleParams, status int) (int, error) {
	articleModel := model.ArticleModelInstance()
	// fixme: author id needed
	return articleModel.AddArticle(params.Title, params.Content, 0, status)
}
