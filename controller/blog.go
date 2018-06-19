package controller

import (
	"net/http"
	"github.com/Kongoole/minreuse-go/model"
)

type Blog struct {}

// Blog shows blog list
func (b Blog) Index(w http.ResponseWriter, r *http.Request) {
	model.ArticleModel{}.FetchAll()
	w.Write([]byte("blog index page"))
}

// Article shows an article
func (b Blog) Article(w http.ResponseWriter, r *http.Request) {

}


