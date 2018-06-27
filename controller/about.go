package controller

import (
	"net/http"

	"github.com/kongoole/minreuse-go/render"
)

type About struct {
	Controller
}

func (a About) Index(w http.ResponseWriter, r *http.Request) {
	render.New().SetDestination(w).SetHasSlogan(false).SetTemplates("about.html").View(nil)
}
