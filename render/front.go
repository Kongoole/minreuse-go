package render

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type FrontRender struct {
	templates []string
	hasHeader bool
	hasFooter bool
	hasSlogan bool
	hasTags   bool
}

func NewFrontRender() *FrontRender {
	return &FrontRender{
		hasHeader: true,
		hasFooter: true,
		hasSlogan: true,
		hasTags:   true,
	}
}

func (r *FrontRender) SetHasHeader(status bool) *FrontRender {
	r.hasHeader = status
	return r
}

func (r *FrontRender) SetHasFooter(status bool) *FrontRender {
	r.hasFooter = status
	return r
}

func (r *FrontRender) SetHasSlogan(status bool) *FrontRender {
	r.hasSlogan = status
	return r
}

func (r *FrontRender) SetHasTags(status bool) *FrontRender {
	r.hasTags = status
	return r
}

func (r *FrontRender) SetTemplates(fileNames ...string) *FrontRender {
	viewFolder := os.Getenv("view_folder")
	for _, fileName := range fileNames {
		r.templates = append(r.templates, viewFolder+fileName)
	}
	return r
}

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.Run([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}

func (r *FrontRender) Render(w io.Writer, data interface{}) {
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
	t.Execute(w, data)
}
