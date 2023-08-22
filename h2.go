package carbon

import "io"

type h2 struct {
	children []Component
	attrs    []Attr
}

var _ Component = (*h2)(nil)

func H2(c ...Component) *h2 {
	return &h2{
		children: c,
		attrs:    nil,
	}
}

func (h *h2) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *h2) Render(w io.Writer) {
	w.Write([]byte("<h2"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	renderAll(h.children, w)
	w.Write([]byte("</h2>"))
}
