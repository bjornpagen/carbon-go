package carbon

import "io"

type content struct {
	children any
	attrs    []Attr
}

var _ Component = (*content)(nil)

func Content(a ...any) *content {
	return &content{
		children: a,
		attrs:    nil,
	}
}

func (c *content) Attr(name string, value string) *content {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *content) Render(w io.Writer) {
	w.Write([]byte(`<main id="main-content" class="bx--content">`))
	renderAny(w, c.children)
	w.Write([]byte(`</main>`))
}
