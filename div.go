package carbon

import (
	"io"
)

type div struct {
	children any
	attrs    []Attr
}

var _ Component = (*div)(nil)

func Div(a ...any) *div {
	return &div{
		children: a,
		attrs:    nil,
	}
}

func (d *div) Attr(name string, value string) Component {
	d.attrs = append(d.attrs, Attr{name, value})
	return d
}

func (d *div) Render(w io.Writer) {
	w.Write([]byte(`<div`))
	renderAttrs(w, d.attrs)
	w.Write([]byte(`>`))
	renderAny(w, d.children)
	w.Write([]byte(`</div>`))
}
