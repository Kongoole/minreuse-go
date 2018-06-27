package controller

import (
	"net/http"

	"github.com/kongoole/minreuse-go/render"
)

type Admin struct {
	Controller
}

func (a Admin) Index(w http.ResponseWriter, r *http.Request) {
	render.New().SetDestination(w).SetTemplates("admin/index.html").View(nil)
}
