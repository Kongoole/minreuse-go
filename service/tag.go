package service

import (
	"github.com/pkg/errors"
	"github.com/kongoole/minreuse-go/model"
)

type TagService struct {}
type ArticleTagService struct{
	ArticleTagModel model.ArticleTagModel
}

func NewTagService() *TagService {
	return &TagService{}
}
func NewArticleTagService(articleTagModel model.ArticleTagModel) * ArticleTagService {
	return &ArticleTagService{ArticleTagModel:articleTagModel}
}

func (ats *ArticleTagService) AddArticleTags(articleID int, tagIDs []int) (int64, error) {
	if len(tagIDs) == 0 {
		return 0, errors.New("tag ids cannot be empty")
	}
	if articleID == 0 {
		return 0, errors.New("invalid article id")
	}
	return ats.ArticleTagModel.AddArticleTags(articleID, tagIDs)
}