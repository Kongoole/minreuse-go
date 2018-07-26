package service

import (
	"github.com/kongoole/minreuse-go/model"
)

type ArticleService struct{}

func ArticleServiceInstance() *ArticleService {
	return new(ArticleService)
}

func (as *ArticleService) UpdateArticle(articleId int, data map[string]interface{}) {
	model.ArticleModelInstance().UpdateArticle(articleId, data)
}
