package controller

import (
	"net/http"
	"html/template"
)

type Home struct {}

// home shows home page
func (h Home) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/base.html")
	t.Execute(w, nil)
}