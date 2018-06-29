package render

import (
	"html/template"
	"io"
	"os"
	"strings"
)

type AdminRender struct {
	templates []string
}

func NewAdminRender() *AdminRender {
	return &AdminRender{}
}

func (a *AdminRender) SetTemplates(fileNames ...string) *AdminRender {
	viewFolder := os.Getenv("view_folder")
	for _, fileName := range fileNames {
		a.templates = append(a.templates, viewFolder+fileName)
	}
	a.templates = append(a.templates, viewFolder+"admin/header.html", viewFolder+"admin/footer.html")
	return a
}

func (a *AdminRender) Render(w io.Writer, data interface{}) {
	// New() must has a parameter, here use the first file name
	t, _ := template.New(strings.Split(a.templates[0], "/")[2]).Funcs(template.FuncMap{"html": unescape}).ParseFiles(a.templates...)
	t.Execute(w, data)
}
