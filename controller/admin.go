package controller

import (
	"encoding/json"
	"github.com/kongoole/minreuse-go/utils/log"
	"net/http"
	"strconv"

	"github.com/kongoole/minreuse-go/model"
	"github.com/kongoole/minreuse-go/render"
	"github.com/kongoole/minreuse-go/service"
)

type Admin struct {
	Controller
}

func (a Admin) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data map[string]interface{}
		json.NewDecoder(r.Body).Decode(&data)
		loginService := service.LoginService()
		isPwdValid := loginService.CheckLogin(data["account"].(string), data["pwd"].(string))
		if isPwdValid {
			// set session
			s := service.NewSessionManager()
			session, err := s.Store.Get(r, s.DefaultSessionName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			session.Values["is_login"] = 1
			session.Save(r, w)
			resp := service.Response{Code: http.StatusOK, Msg: "success", Data: nil}
			resp.JSONResponse(w)
			return
		}
		resp := service.Response{Code: http.StatusBadRequest, Msg: "fail", Data: nil}
		resp.JSONResponse(w)
	} else {
		render.NewAdminRender().SetTemplates("admin/login.html").Render(w, nil)
	}
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
			log.Fatal("fail to get off")
		}
	}
	articleModel := model.NewArticleModel()
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
	params := service.AddArticleParams{}
	json.NewDecoder(r.Body).Decode(&params)
	articleService := service.NewArticleService()
	articleService.AddArticle(params, model.ArticleModel{}.StatusUnpublished)
}

func (a Admin) PublishArticle(w http.ResponseWriter, r *http.Request) {
	params := service.AddArticleParams{}
	json.NewDecoder(r.Body).Decode(&params)
	articleService := service.NewArticleService()
	articleModel := model.NewArticleModel()
	articleID, err := articleService.AddArticle(params, articleModel.StatusPublished)
	if err != nil {
		log.Debug("fail to add article, error msg:" + err.Error())
	}
	articleTagService := service.NewArticleTagService(model.ArticleTagModel{})
	_, err = articleTagService.AddArticleTags(articleID, params.TagIds)
	if err != nil {
		log.Debug("fail to add article tags with error msg: " + err.Error())
	}
}

// EditArticle shows article edit page
func (a Admin) EditArticle(w http.ResponseWriter, r *http.Request) {
	articleId, _ := strconv.Atoi(r.URL.Query().Get("article_id"))
	article := model.NewArticleModel().FetchOneByArticleId(articleId)
	tags := model.NewTagModel().FetchTagsByArticleId(articleId)
	allTags := model.NewTagModel().FetchAll()
	data := struct {
		Article model.Article
		Tags    []model.Tag
		AllTags []model.Tag
	}{article, tags, allTags}
	render.NewAdminRender().SetTemplates("admin/article_edit.html").Render(w, data)
}

// UpdateArticle updates an article
func (a Admin) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)
	articleID, _ := strconv.Atoi(data["article_id"].(string))
	delete(data, "article_id")
	delete(data, "tagIds")
	updated := service.NewArticleService().UpdateArticle(articleID, data)
	if !updated {
		service.Response{http.StatusBadGateway, "fail to update", nil}.JSONResponse(w)
	} else {
		service.Response{http.StatusOK, "success", nil}.JSONResponse(w)
	}
}
