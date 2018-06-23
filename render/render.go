package render

import (
	"os"
	"html/template"
	"net/http"
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"strings"
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

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.Run([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
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
	// New() must has a parameter, here use the first file name
	t, _ := template.New(strings.Split(r.templates[0], "/")[1]).Funcs(template.FuncMap{"markDown": markDowner}).ParseFiles(r.templates...)
	t.Execute(r.wr, data)
}
