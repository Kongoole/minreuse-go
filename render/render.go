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
	hasSlogan bool
	hasTags   bool
}

func New() *Render {
	return &Render{
		hasHeader: true,
		hasFooter: true,
		hasSlogan: true,
		hasTags:   true,
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

func (r *Render) SetHasSlogan(status bool) *Render {
	r.hasSlogan = status
	return r
}

func (r *Render) SetHasTags(status bool) *Render {
	r.hasTags = status
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
	if r.hasSlogan {
		r.templates = append(r.templates, os.Getenv("view_folder")+"common/slogan.html")
	}
	if r.hasTags {
		r.templates = append(r.templates, os.Getenv("view_folder")+"common/tag.html")
	}
	t, _ := template.ParseFiles(r.templates...)
	t.Execute(r.wr, data)
}
