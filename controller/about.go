package controller

import (
	"net/http"

	"github.com/kongoole/minreuse-go/render"
)

type About struct {
	Controller
}

func (a About) Index(w http.ResponseWriter, r *http.Request) {
	render.NewFrontRender().SetHasSlogan(false).SetTemplates("about.html").Render(w, nil)
}
