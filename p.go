package carbon

import (
	"io"
)

type p struct {
	children any
	attrs    []Attr
}

var _ Component = (*p)(nil)

func P(a ...any) *p {
	return &p{
		children: a,
		attrs:    nil,
	}
}

func (d *p) Attr(name string, value string) Component {
	d.attrs = append(d.attrs, Attr{name, value})
	return d
}

func (d *p) Render(w io.Writer) {
	w.Write([]byte(`<p`))
	renderAttrs(w, d.attrs)
	w.Write([]byte(`>`))
	renderAny(w, d.children)
	w.Write([]byte(`</p>`))
}
