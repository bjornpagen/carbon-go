package carbon

import "io"

type h4 struct {
	children []Component
	attrs    []Attr
}

var _ Component = (*h4)(nil)

func H4(c ...Component) *h4 {
	return &h4{
		children: c,
		attrs:    nil,
	}
}

func (h *h4) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *h4) Render(w io.Writer) {
	w.Write([]byte("<h4"))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	renderAll(h.children, w)
	w.Write([]byte("</h4>"))
}
