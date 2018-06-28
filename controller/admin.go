package controller

import (
	"net/http"

	"github.com/kongoole/minreuse-go/render"
	"github.com/kongoole/minreuse-go/model"
)

type Admin struct {
	Controller
}

func (a Admin) Index(w http.ResponseWriter, r *http.Request) {
	render.NewAdminRender().SetTemplates("admin/index.html").Render(w, nil)
}

func (a Admin) ArticleList(w http.ResponseWriter, r *http.Request) {
	articles := model.NewArticleModel().FetchAll()
	render.NewAdminRender().SetTemplates("admin/article_list.html").Render(w, articles)
}

func (a Admin) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	render.NewAdminRender().SetTemplates("admin/article_create.html").Render(w, nil)
}
