package route

import (
	"net/http"

	"github.com/kongoole/minreuse-go/service"

	"github.com/kongoole/minreuse-go/controller"
)

func Register() {
	http.HandleFunc("/", controller.Home{}.Index)
	http.HandleFunc("/blog", controller.Blog{}.Index)
	http.HandleFunc("/article", controller.Blog{}.Article)
	http.HandleFunc("/tag", controller.Blog{}.TagArticles)
	http.HandleFunc("/blog/search", controller.Blog{}.Search)
	http.HandleFunc("/about", controller.About{}.Index)

	http.HandleFunc("/admin", controller.Admin{}.Login)
	http.HandleFunc("/admin/login", controller.Admin{}.Login)
	http.HandleFunc("/admin/index", AuthMiddleWare(controller.Admin{}.Index))
	http.HandleFunc("/admin/article/create", AuthMiddleWare(controller.Admin{}.ArticleCreate))
	http.HandleFunc("/admin/article/list", AuthMiddleWare(controller.Admin{}.ArticleList))
	http.HandleFunc("/admin/article/save", AuthMiddleWare(controller.Admin{}.SaveArticle))
	http.HandleFunc("/admin/article/publish", AuthMiddleWare(controller.Admin{}.PublishArticle))
	http.HandleFunc("/admin/article/edit", AuthMiddleWare(controller.Admin{}.EditArticle))
	http.HandleFunc("/admin/article/update", AuthMiddleWare(controller.Admin{}.UpdateArticle))

	http.HandleFunc("/public/", serveResource)
}

// serveResource returns static files
func serveResource(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
}

func AuthMiddleWare(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sm := service.NewSessionManager()
		s, _ := sm.Store.Get(r, sm.DefaultSessionName)
		if s.Values["is_login"] == 1 {
			f.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
