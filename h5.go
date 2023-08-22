package carbon

import "io"

type h5 struct {
	children []Component
	attrs    []Attr
}

var _ Component = (*h5)(nil)

func H5(c ...Component) *h5 {
	return &h5{
		children: c,
		attrs:    nil,
	}
}

func (h *h5) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *h5) Render(w io.Writer) {
	w.Write([]byte("<h5"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	renderAll(h.children, w)
	w.Write([]byte("</h5>"))
}
