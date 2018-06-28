package service

import (
	"strconv"

	"github.com/kongoole/minreuse-go/model"
)

type pagination struct {
}

func NewPagination() pagination {
	return pagination{}
}

// Html makes pagination html
func (p pagination) Html(total int, offset int) string {
	// caculate pages
	var pages int
	if total == model.PAGE_SIZE {
		pages = 1
	} else {
		pages = total/model.PAGE_SIZE + 1
	}
	// generate pagination html
	var pagination string
	if pages > 1 {
		for i := 1; i <= pages; i++ {
			class := "waves-effect"
			if i == offset+1 {
				class = "active"
			}
			pagination = pagination + "<li class=\"" + class + "\"><a href=\"/blog?page=" + strconv.Itoa(i-1) + "\">" +
				strconv.Itoa(i) + "</a></li>"
		}
	}
	return pagination
}
