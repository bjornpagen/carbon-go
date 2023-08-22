package carbon

import "io"

type h3 struct {
	children []Component
	attrs    []Attr
}

var _ Component = (*h3)(nil)

func H3(c ...Component) *h3 {
	return &h3{
		children: c,
		attrs:    nil,
	}
}

func (h *h3) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *h3) Render(w io.Writer) {
	w.Write([]byte("<h3"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	renderAll(h.children, w)
	w.Write([]byte("</h3>"))
}
