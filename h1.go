package carbon

import "io"

type h1 struct {
	children []Component
	attrs    []Attr
}

var _ Component = (*h1)(nil)

func H1(c ...Component) *h1 {
	return &h1{
		children: c,
		attrs:    nil,
	}
}

func (h *h1) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *h1) Render(w io.Writer) {
	w.Write([]byte("<h1"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	renderAll(h.children, w)
	w.Write([]byte("</h1>"))
}
