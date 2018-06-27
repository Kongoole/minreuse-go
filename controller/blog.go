package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kongoole/minreuse-go/model"
	"github.com/kongoole/minreuse-go/render"
	"github.com/kongoole/minreuse-go/service"
)

type Blog struct {
	Controller
}

// BlogData is used to render blog page
type BlogData struct {
	Articles []model.Article
	Tags     []model.Tag
	Keywords string
}

// Blog shows blog list
func (b Blog) Index(w http.ResponseWriter, r *http.Request) {
	articles := model.NewArticleModel().FetchAll()
	tags := model.NewTagModel().FetchTagsWithArticlesNum()
	data := BlogData{Articles: articles, Tags: tags}
	render.NewFrontRender().SetTemplates("blog.html").Render(w, data)
}

// Article shows an article
func (b Blog) Article(w http.ResponseWriter, r *http.Request) {
	articleId, err := strconv.Atoi(r.URL.Query().Get("article_id"))
	if err != nil {
		log.Fatal(err)
	}
	article := model.NewArticleModel().FetchOneByArticleId(articleId)
	tags := model.NewTagModel().FetchTagsByArticleId(articleId)
	data := struct {
		Article model.Article
		Tags    []model.Tag
	}{article, tags}
	render.NewFrontRender().SetTemplates("article.html").SetHasSlogan(false).Render(w, data)
}

func (b Blog) TagArticles(w http.ResponseWriter, r *http.Request) {
	tagId, err := strconv.Atoi(r.URL.Query().Get("tag_id"))
	if err != nil {
		log.Fatal(err)
	}
	articles := model.NewArticleModel().FetchTagArticlesByTagId(tagId)
	tags := model.NewTagModel().FetchTagsWithArticlesNum()
	data := BlogData{Articles: articles, Tags: tags}
	render.NewFrontRender().SetTemplates("blog.html").Render(w, data)
}

func (b Blog) Search(w http.ResponseWriter, r *http.Request) {
	searcher := service.NewArticleSearcher()
	keywords := r.URL.Query().Get("keywords")
	articles := service.DoSearch(searcher, keywords).([]model.Article)
	tags := model.NewTagModel().FetchTagsWithArticlesNum()
	data := BlogData{articles, tags, keywords}
	render.NewFrontRender().SetTemplates("blog.html").Render(w, data)
}
