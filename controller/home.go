package controller

import (
	"net/http"
	"github.com/kongoole/minreuse-go/render"
)

type Home struct {}

// Index shows home page
func (h Home) Index(w http.ResponseWriter, r *http.Request) {
	render.New().SetDestination(w).SetTemplates("home.html").View(nil)
}