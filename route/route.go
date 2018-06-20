package route

import (
	"net/http"
	"github.com/Kongoole/minreuse-go/controller"
)

func Register() {
	http.HandleFunc("/", controller.Home{}.Index)
	http.HandleFunc("/blog", controller.Blog{}.Index)
	http.HandleFunc("/article", controller.Blog{}.Article)
	http.HandleFunc("/public/", serveResource)
}

// serveResource returns static files
func serveResource(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
}
