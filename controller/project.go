package controller

import (
	"net/http"
	"github.com/kongoole/minreuse-go/render"
)

type Project struct {
	Controller
}

func (p Project) Index(w http.ResponseWriter, r *http.Request) {
	render.NewFrontRender().SetHasSlogan(false).SetTemplates("project.html").Render(w, nil)
}
