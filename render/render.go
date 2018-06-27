package render

import (
	"io"
)

type Render interface {
	Renderer(w io.Writer, data interface{})
}
