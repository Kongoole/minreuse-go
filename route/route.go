package route

import (
	"net/http"

	"github.com/kongoole/minreuse-go/controller"
)

func Register() {
	http.HandleFunc("/", controller.Home{}.Index)
	http.HandleFunc("/blog", controller.Blog{}.Index)
	http.HandleFunc("/article", controller.Blog{}.Article)
	http.HandleFunc("/tag", controller.Blog{}.TagArticles)
	http.HandleFunc("/blog/search", controller.Blog{}.Search)
	http.HandleFunc("/about", controller.About{}.Index)

	http.HandleFunc("/admin/index", controller.Admin{}.Index)
	http.HandleFunc("/admin/article/create", controller.Admin{}.ArticleCreate)
	http.HandleFunc("/admin/article/list", controller.Admin{}.ArticleList)
	http.HandleFunc("/admin/article/save", controller.Admin{}.SaveArticle)
	http.HandleFunc("/admin/article/publish", controller.Admin{}.PublishArticle)
	http.HandleFunc("/admin/article/edit", controller.Admin{}.EditArticle)
	http.HandleFunc("/admin/article/update", controller.Admin{}.UpdateArticle)

	http.HandleFunc("/public/", serveResource)
}

// serveResource returns static files
func serveResource(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
}
