package carbon

import "io"

type basic struct {
	_tag     string
	children []Component
	attrs    []Attr
}

var _ Component = (*basic)(nil)

func (h *basic) Attr(name string, value string) Component {
	h.attrs = append(h.attrs, Attr{name, value})
	return h
}

func (h *basic) Render(w io.Writer) {
	w.Write([]byte("<"))
	w.Write(yoloBytesUnsafe(h._tag))
	renderAttrs(w, h.attrs)
	w.Write([]byte(">"))
	for _, c := range h.children {
		c.Render(w)
	}
	w.Write([]byte("</"))
	w.Write(yoloBytesUnsafe(h._tag))
	w.Write([]byte(">"))
}

func P(c ...Component) *basic {
	return &basic{
		_tag:     "p",
		children: c,
		attrs:    nil,
	}
}

func H1(c ...Component) *basic {
	return &basic{
		_tag:     "h1",
		children: c,
		attrs:    nil,
	}
}

func H2(c ...Component) *basic {
	return &basic{
		_tag:     "h2",
		children: c,
		attrs:    nil,
	}
}

func H3(c ...Component) *basic {
	return &basic{
		_tag:     "h3",
		children: c,
		attrs:    nil,
	}
}

func H4(c ...Component) *basic {
	return &basic{
		_tag:     "h4",
		children: c,
		attrs:    nil,
	}
}

func H5(c ...Component) *basic {
	return &basic{
		_tag:     "h5",
		children: c,
		attrs:    nil,
	}
}

func H6(c ...Component) *basic {
	return &basic{
		_tag:     "h6",
		children: c,
		attrs:    nil,
	}
}

func Div(c ...Component) *basic {
	return &basic{
		_tag:     "div",
		children: c,
		attrs:    nil,
	}
}

func Span(c ...Component) *basic {
	return &basic{
		_tag:     "span",
		children: c,
		attrs:    nil,
	}
}

func Em(c ...Component) *basic {
	return &basic{
		_tag:     "em",
		children: c,
		attrs:    nil,
	}
}

func Strong(c ...Component) *basic {
	return &basic{
		_tag:     "strong",
		children: c,
		attrs:    nil,
	}
}
