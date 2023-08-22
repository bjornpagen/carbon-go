package carbon

import "io"

type h6 struct {
	children []Component
	attrs    []Attr
}

var _ Component = (*h6)(nil)

func H6(c ...Component) *h6 {
	return &h6{
		children: c,
		attrs:    nil,
	}
}

func (h *h6) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *h6) Render(w io.Writer) {
	w.Write([]byte("<h6"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	renderAll(h.children, w)
	w.Write([]byte("</h6>"))
}
