package carbon

import "io"

type content struct {
	children any
	attrs    []Attr

	id string
}

var _ Component = (*content)(nil)

func Content(a ...any) *content {
	return &content{
		children: a,
		attrs:    nil,

		id: "main-content",
	}
}

func (c *content) Attr(name string, value string) Component {
	c.attrs = append(c.attrs, Attr{name, value})
	return c
}

func (c *content) Id(id string) *content {
	c.id = id
	return c
}

func (c *content) Render(w io.Writer) {
	w.Write([]byte(`<main id="`))
	io.WriteString(w, c.id)
	w.Write([]byte(`" class="bx--content">`))
	renderAny(w, c.children)
	w.Write([]byte(`</main>`))
}
