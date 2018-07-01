package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kongoole/minreuse-go/model"
	"github.com/kongoole/minreuse-go/render"
	"github.com/kongoole/minreuse-go/service"
)

type Admin struct {
	Controller
}

func (a Admin) Index(w http.ResponseWriter, r *http.Request) {
	render.NewAdminRender().SetTemplates("admin/index.html").Render(w, nil)
}

func (a Admin) ArticleList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	offset := 0
	if page != "" {
		var err error
		offset, err = strconv.Atoi(page)
		if err != nil {
			log.Println("fail to get off")
		}
	}
	articleModel := model.NewArticleModel()
	articles := articleModel.FetchWithPagination(offset)
	total := articleModel.FetchArticleAmount()
	pagination := service.NewPagination().HTML(total, offset, "/admin/article/list")
	data := struct {
		Articles   []model.Article
		Pagination string
	}{articles, pagination}
	render.NewAdminRender().SetTemplates("admin/article_list.html").Render(w, data)
}

func (a Admin) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	tags := model.NewTagModel().FetchAll()
	data := struct {
		Tags []model.Tag
	}{tags}
	render.NewAdminRender().SetTemplates("admin/article_create.html").Render(w, data)
}
