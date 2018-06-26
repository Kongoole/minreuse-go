package service

import (
	"github.com/kongoole/minreuse-go/model"
)

type Searcher interface {
	Search(keywords string) interface{}
}

type ArticleSearcher struct {
	model model.ArticleModel
}

// DOSearch() does search action
func DoSearch(s Searcher, keywords string) interface{} {
	return s.Search(keywords)
}

func NewArticleSearcher() ArticleSearcher {
	return ArticleSearcher{}
}

func (as ArticleSearcher) Search(keywords string) interface{} {
	return as.model.FetchArticlesByKeyWords(keywords)
}
