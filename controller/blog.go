package controller

import (
	"net/http"
	"github.com/Kongoole/minreuse-go/model"
	"github.com/Kongoole/minreuse-go/render"
)

type Blog struct{}

// Blog shows blog list
func (b Blog) Index(w http.ResponseWriter, r *http.Request) {
	articles := model.ArticleModel{}.FetchAll()
	render.New().SetDestination(w).SetTemplates("blog.html").View(articles)
}

// Article shows an article
func (b Blog) Article(w http.ResponseWriter, r *http.Request) {

	render.New().SetDestination(w).SetTemplates("article.html").SetHasSlogan(false).View(nil)
}
