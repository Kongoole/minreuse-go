package controller

import (
	"net/http"

	"github.com/kongoole/minreuse-go/render"
)

type Admin struct {
	Controller
}

func (a Admin) Index(w http.ResponseWriter, r *http.Request) {
	render.NewAdminRender().SetTemplates("admin/index.html").Render(w, nil)
}

func (a Admin) CreateArticle(w http.ResponseWriter, r *http.Request) {
	render.NewAdminRender().SetTemplates("admin/article.html").Render(w, nil)
}
