package render

import (
	"os"
	"text/template"
	"net/http"
)

type Render struct {
	templates []string
	wr        http.ResponseWriter
	hasHeader bool
	hasFooter bool
}

func New() *Render {
	return &Render{
		hasHeader: true,
		hasFooter: true,
	}
}

func (r *Render) SetHasHeader(status bool) *Render {
	r.hasHeader = status
	return r
}

func (r *Render) SetHasFooter(status bool) *Render {
	r.hasFooter = status
	return r
}

func (r *Render) SetTemplates(fileNames ...string) *Render {
	viewFolder := os.Getenv("view_folder")
	for _, fileName := range fileNames {
		r.templates = append(r.templates, viewFolder+fileName)
	}
	return r
}

func (r *Render) SetDestination(wr http.ResponseWriter) *Render {
	r.wr = wr
	return r
}

func (r *Render) View(data interface{}) {
	if r.hasHeader {
		r.templates = append(r.templates, os.Getenv("view_folder")+"common/header.html")
	}
	if r.hasFooter {
		r.templates = append(r.templates, os.Getenv("view_folder")+"common/footer.html")
	}
	t, _ := template.ParseFiles(r.templates...)
	t.Execute(r.wr, data)
}
