package controller

import (
	"encoding/json"
	"fmt"
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
	articleModel := model.ArticleModelInstance()
	articles := articleModel.FetchWithPagination(offset, articleModel.StatusPublished, articleModel.StatusUnpublished)
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

func (a Admin) SaveArticle(w http.ResponseWriter, r *http.Request) {
	addArticle(w, r, model.ArticleModelInstance().StatusUnpublished)
}

func (a Admin) PublishArticle(w http.ResponseWriter, r *http.Request) {
	addArticle(w, r, model.ArticleModelInstance().StatusPublished)
}

func addArticle(w http.ResponseWriter, r *http.Request, status int) {
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	// tagIds := r.FormValue("tagIds")
	articleModel := model.ArticleModelInstance()
	_, err := articleModel.AddArticle(title, content, 0, status)
	if err != nil {
		resp, _ := json.Marshal(service.Response{Code: service.HTTP_SERVER_ERROR, Msg: err.Error(), Data: nil})
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(service.Response{Code: service.HTTP_OK, Msg: "success", Data: nil})
	w.Write(resp)
}

func (a Admin) EditArticle(w http.ResponseWriter, r *http.Request) {

}

func (a Admin) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Println(title, content)
}
