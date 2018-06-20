package controller

import (
	"net/http"
	"github.com/Kongoole/minreuse-go/render"
)

type Home struct {}

// home shows home page
func (h Home) Index(w http.ResponseWriter, r *http.Request) {
	render.New().SetDestination(w).SetTemplates("home.html").View()
}