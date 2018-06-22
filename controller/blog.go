package controller

import (
	"net/http"
	"github.com/Kongoole/minreuse-go/model"
	"github.com/Kongoole/minreuse-go/render"
	"strconv"
	"log"
)

type Blog struct{}

// Blog shows blog list
func (b Blog) Index(w http.ResponseWriter, r *http.Request) {
	articles := model.ArticleModel{}.FetchAll()
	render.New().SetDestination(w).SetTemplates("blog.html").View(articles)
}

// Article shows an article
func (b Blog) Article(w http.ResponseWriter, r *http.Request) {
	articleId, err := strconv.Atoi(r.URL.Query().Get("article_id"))
	if err != nil {
		log.Fatal(err)
	}
	article := model.ArticleModel{}.FetchOneByArticleId(articleId)
	tags := model.TagModel{}.FetchTagsByArticleId(articleId)
	data := struct {
		Article model.Article
		Tags []model.Tag
	}{article, tags}
	render.New().SetDestination(w).SetTemplates("article.html").SetHasSlogan(false).View(data)
}
